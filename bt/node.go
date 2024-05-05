package bt

type node[T any] struct {
	value      T
	leftChild  *node[T]
	rightChild *node[T]
}

type balanceableNode[T any] struct {
	value      T
	leftChild  *balanceableNode[T]
	rightChild *balanceableNode[T]
	height     int
}

func (n *balanceableNode[T]) setHeight() {
	n.height = 1 + max(n.leftChild.getHeight(), n.rightChild.getHeight())
}

func (n *balanceableNode[T]) getHeight() int {
	if n == nil {
		return -1
	}
	return n.height
}

func (n *balanceableNode[T]) rotateLeft() *balanceableNode[T] {
	newRoot := n.rightChild
	newRoot.leftChild, n.rightChild = n, newRoot.leftChild
	n.setHeight()
	newRoot.setHeight()
	return newRoot
}

func (n *balanceableNode[T]) rotateLeftRight() *balanceableNode[T] {
	n.leftChild = n.leftChild.rotateLeft()
	return n.rotateRight()
}

func (n *balanceableNode[T]) rotateRight() *balanceableNode[T] {
	newRoot := n.leftChild
	newRoot.rightChild, n.leftChild = n, newRoot.rightChild
	n.setHeight()
	newRoot.setHeight()
	return newRoot
}

func (n *balanceableNode[T]) rotateRightLeft() *balanceableNode[T] {
	n.rightChild = n.rightChild.rotateRight()
	return n.rotateLeft()
}
