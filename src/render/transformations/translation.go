package transformations

import (
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

func Translate(v Vec4, t Mat4) Vec4 {
	return t.Mul4x1(v)
}
