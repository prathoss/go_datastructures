package xmath

type Comp uint8

const (
	// CompEq the elements are equal
	CompEq Comp = iota
	// CompGt the first element is greater than the second one
	CompGt
	// CompLs the first element is lesser than the second one
	CompLs
)
