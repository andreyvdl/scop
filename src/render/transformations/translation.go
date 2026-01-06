package transformations

import (
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

func Translate(trans Mat4, t Vec3) Mat4 {
	id := Identity4()

	id[0][3] = t[0]
	id[1][3] = t[1]
	id[2][3] = t[2]

	return trans.Mul(id)
}
