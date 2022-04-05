package xmath

type Comp uint8

const (
	// CompEq the elements are equal
	CompEq Comp = 1
	// CompGt the first element is greater than the second one
	CompGt Comp = 2
	// CompLs the first element is lesser than the second one
	CompLs Comp = 3
)
