package transformations

import (
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

func Combine(v Vec4, t ...Mat4) Vec4 {
	res := Identity4()

	for i := len(t) - 1; i >= 0; i-- {
		res = res.Mul(t[i])
	}

	return res.Mul4x1(v)
}
