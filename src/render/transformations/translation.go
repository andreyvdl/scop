package transformations

import (
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

func Translate(t Vec3) Mat4 {
	id := Identity4()

	id[3] = t[0]
	id[7] = t[1]
	id[11] = t[2]
	return id
}
