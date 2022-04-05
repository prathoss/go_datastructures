package bt

type node[T any] struct {
	value      T
	leftChild  *node[T]
	rightChild *node[T]
}

type nodeP[T any] struct {
	value      T
	leftChild  *nodeP[T]
	rightChild *nodeP[T]
	parent     *nodeP[T]
}
