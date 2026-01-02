package transformations

import (
	"math"
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

func Rotate(theta float64, axis Vec3) Mat4 {
	id := Identity4()
	cos := float32(math.Cos(theta))
	sin := float32(math.Sin(theta))

	id[0] = cos + axis[0]*axis[0]*(1-cos)
	id[1] = axis[0]*axis[1]*(1-cos) - axis[2]*sin
	id[2] = axis[0]*axis[2]*(1-cos) + axis[1]*sin

	id[4] = axis[1]*axis[0]*(1-cos) + axis[2]*sin
	id[5] = cos + axis[1]*axis[1]*(1-cos)
	id[6] = axis[1]*axis[2]*(1-cos) - axis[0]*sin

	id[8] = axis[2]*axis[0]*(1-cos) - axis[1]*sin
	id[9] = axis[2]*axis[1]*(1-cos) + axis[0]*sin
	id[10] = cos + axis[2]*axis[2]*(1-cos)

	return id
}
