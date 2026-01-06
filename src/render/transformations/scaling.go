package transformations

import (
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

func Scale(trans Mat4, s Vec3) Mat4 {
	id := Identity4()

	id[0][0] = s[0]
	id[1][1] = s[1]
	id[2][2] = s[2]

	return trans.Mul(id)
}
