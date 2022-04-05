package bt

import (
	"math"

	"github.com/prathoss/gods/xmath"
)

// NewBinaryTree creates binary tree with comparing function which will decide the value of node
// you can use just CpFunc for basic ordered types or wrap the CpFunc:
//
//  type myStruct struct {
//  	comparableField int
//  }
//  bt := NewBinaryTree(func(a, b myStruct) xmath.Comp {
//  	return CpFunc(a.comparableField, b.comparableField)
//  }
func NewBinaryTree[T any](cpFunc func(a, b T) xmath.Comp) *BinaryTree[T] {
	return &BinaryTree[T]{cpFunc: cpFunc}
}

type BinaryTree[T any] struct {
	root   *node[T]
	cpFunc func(a, b T) xmath.Comp
}

func (b *BinaryTree[T]) Add(element T) {
	n := &node[T]{value: element}
	if b.root == nil {
		b.root = n
		return
	}
	b.addNotRoot(n, b.root)
}

func (b *BinaryTree[T]) addNotRoot(new, curr *node[T]) {
	cp := b.cpFunc(new.value, curr.value)
	switch cp {
	case xmath.CompLs:
		if curr.leftChild == nil {
			curr.leftChild = new
		} else {
			b.addNotRoot(new, curr.leftChild)
		}
	case xmath.CompGt:
		if curr.rightChild == nil {
			curr.rightChild = new
		} else {
			b.addNotRoot(new, curr.rightChild)
		}
	default:
		// handle as if the two values were equal (binary tree can contain each value only once) => do not add
		return
	}
}

func (b BinaryTree[T]) Find(element T) (bool, T) {
	return b.find(element, b.root)
}

func (b BinaryTree[T]) find(element T, currNode *node[T]) (bool, T) {
	if currNode == nil {
		var empty T
		return false, empty
	}
	cp := b.cpFunc(element, currNode.value)
	switch cp {
	case xmath.CompLs:
		return b.find(element, currNode.leftChild)
	case xmath.CompGt:
		return b.find(element, currNode.rightChild)
	default:
		return true, currNode.value
	}
}

// Print returns 2D slice of elements separated by tabulator:
//
//  [	,2	,	]
//  [1	,	,3	]
func (b BinaryTree[T]) Print(nodePrinter func(ele T) string) [][]string {
	// maximal number of nodes at a level is 2^n where n is a height - 1
	height := b.getHeight()
	if height == 0 {
		return [][]string{}
	}
	maxNumberOfNodes := 0
	for i := 0; i < height; i++ {
		maxNumberOfNodes += getNumberOfNodesInLevel(i)
	}

	/*
		separate the line to sections of not printed fields, the number of separators is the number of maximal nodes of a tree of current height,
		skip every second separator (it is separator from previous level, fill it with empty "tab")
		|0|1|2|3|4|5|6| slice indexes
		| | | |3| | | | (2 sections, 1 printed separator)
		| |1| |*| |5| | (4 sections, 1. separator printed, 2. is from first level, 3. printed)
		|0|*|2|*|4|*|6|
	*/
	li := b.GetLevelIterator()
	render := make([][]string, height)
	currentNumberOfNodes := 0
	for i := 0; i < len(render); i++ {
		separatorsVisited := 0
		currentNumberOfNodes += getNumberOfNodesInLevel(i)
		positionOfFirst := maxNumberOfNodes / (currentNumberOfNodes + 1)
		_, nds := li.Next()
		tmp := make([]string, maxNumberOfNodes)
		for j := 0; j < len(tmp); j++ {
			// index of separator
			if j != positionOfFirst+(positionOfFirst+1)*separatorsVisited {
				tmp[j] = "\t"
				continue
			}
			// every second section separator is skipped => it is position of separator from previous level
			if separatorsVisited%2 == 1 {
				tmp[j] = "\t"
				separatorsVisited++
				continue
			}
			// divide separators with 2 to access correct element of nodes returned from iterator
			nd := nds[separatorsVisited/2]
			if nd == nil {
				tmp[j] = "\t"
			} else {
				tmp[j] = nodePrinter(*nd) + "\t"
			}
			separatorsVisited++
		}
		render[i] = tmp
	}
	return render
}

func getNumberOfNodesInLevel(level int) int {
	return int(math.Pow(2, float64(level)))
}

func (b BinaryTree[T]) getHeight() int {
	return getBranchHeight(b.root)
}

func getBranchHeight[T any](branch *node[T]) int {
	if branch == nil {
		return 0
	}
	leftHeight := getBranchHeight(branch.leftChild)
	rightHeight := getBranchHeight(branch.rightChild)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// GetLevelIterator iterates through tree by level, returns slice of elements
// for first level it returns slice with length of 1
func (b BinaryTree[T]) GetLevelIterator() *levelIterator[T] {
	nds := make([]*node[T], 1)
	nds[0] = b.root
	return &levelIterator[T]{
		currentNodes: nds,
	}
}

type levelIterator[T any] struct {
	currentNodes []*node[T]
}

// Next returns bool representing if there was a level to return and slice of nodes at the level
func (l *levelIterator[T]) Next() (bool, []*T) {
	// is this a level or is it just empty (is the result usable)
	shouldContinue := false
	numberOfNodesInLevel := len(l.currentNodes)
	result := make([]*T, numberOfNodesInLevel)
	nxt := make([]*node[T], numberOfNodesInLevel*2)
	for i := 0; i < numberOfNodesInLevel; i++ {
		if l.currentNodes[i] != nil {
			shouldContinue = true
			nxt[2*i] = l.currentNodes[i].leftChild
			nxt[2*i+1] = l.currentNodes[i].rightChild
			result[i] = &l.currentNodes[i].value
		} else {
			nxt[2*i] = nil
			nxt[2*i+1] = nil
			result[i] = nil
		}
	}
	l.currentNodes = nxt
	if shouldContinue {
		return shouldContinue, result
	}
	return shouldContinue, nil
}
