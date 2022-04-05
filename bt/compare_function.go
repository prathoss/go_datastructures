package bt

import (
	"github.com/prathoss/gods/xmath"
	"golang.org/x/exp/constraints"
)

func CpFunc[T constraints.Ordered](a, b T) xmath.Comp {
	if a > b {
		return xmath.CompGt
	} else if a < b {
		return xmath.CompLs
	} else {
		return xmath.CompEq
	}
}
