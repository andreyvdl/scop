package transformations

import (
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

func Scale(v Vec4, s Mat4) Vec4 {
	return s.Mul4x1(v)
}
