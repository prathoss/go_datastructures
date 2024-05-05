package bt

import (
	"github.com/prathoss/gods/xmath"
)

func NewBinarySearchTree[T any](cpFunc func(a, b T) xmath.Comp) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{cpFunc: cpFunc}
}

type BinarySearchTree[T any] struct {
	root   *balanceableNode[T]
	cpFunc func(a, b T) xmath.Comp
}

func (b *BinarySearchTree[T]) Add(element T) {
	n := &balanceableNode[T]{
		value:  element,
		height: 0,
	}
	if b.root == nil {
		b.root = n
		return
	}
	b.addNotRoot(n, b.root)
	b.root = b.balance(b.root)
}

func (b *BinarySearchTree[T]) addNotRoot(new, curr *balanceableNode[T]) {
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

	curr.setHeight()

	// balance children
	if curr.leftChild != nil {
		curr.leftChild = b.balance(curr.leftChild)
	}
	if curr.rightChild != nil {
		curr.rightChild = b.balance(curr.rightChild)
	}
}

func (b *BinarySearchTree[T]) balance(child *balanceableNode[T]) *balanceableNode[T] {
	difference := child.leftChild.getHeight() - child.rightChild.getHeight()
	newSubtreeRoot := child
	// unbalanced right => left rotation
	if difference < -1 {
		if child.rightChild.rightChild == nil ||
			child.rightChild.leftChild.getHeight() > child.rightChild.rightChild.getHeight() {
			newSubtreeRoot = child.rotateRightLeft()
		} else {
			newSubtreeRoot = child.rotateLeft()
		}
	}

	if difference > 1 {
		if child.leftChild == nil ||
			child.leftChild.rightChild.getHeight() > child.leftChild.leftChild.getHeight() {
			newSubtreeRoot = child.rotateLeftRight()
		} else {
			newSubtreeRoot = child.rotateRight()
		}
	}

	return newSubtreeRoot
}

func (b BinarySearchTree[T]) Find(element T) (bool, T) {
	return b.find(element, b.root)
}

func (b BinarySearchTree[T]) find(element T, currNode *balanceableNode[T]) (bool, T) {
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
