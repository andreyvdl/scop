package transformations

import (
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

func Scale(s Vec3) Mat4 {
	id := Identity4()

	id[0] = s[0]
	id[5] = s[1]
	id[10] = s[2]
	return id
}
