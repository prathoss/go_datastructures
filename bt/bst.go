package bt

import "github.com/prathoss/gods/xmath"

func NewBinarySearchTree[T any](cpFunc func(a, b T) xmath.Comp) *BinaryTree[T] {
	return &BinaryTree[T]{cpFunc: cpFunc}
}

type BinarySearchTree[T any] struct {
	root   *nodeP[T]
	cpFunc func(a, b T) xmath.Comp
}

func (b *BinarySearchTree[T]) Add(element T) {
	n := &nodeP[T]{value: element}
	if b.root == nil {
		b.root = n
		return
	}
	b.addNotRoot(n, b.root)
}

func (b *BinarySearchTree[T]) addNotRoot(new, curr *nodeP[T]) {
	new.parent = curr
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

func (b *BinarySearchTree[T]) leftRotation(subtree *nodeP[T]) *nodeP[T] {
	panic("Not implemented!")
}

func (b *BinarySearchTree[T]) rightRotation(subtree *nodeP[T]) *nodeP[T] {
	panic("Not implemented!")
}

func (b BinarySearchTree[T]) Find(element T) (bool, T) {
	return b.find(element, b.root)
}

func (b BinarySearchTree[T]) find(element T, currNode *nodeP[T]) (bool, T) {
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
