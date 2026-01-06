package transformations

import (
	"math"
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

/*
Theta is expected to be in radians
*/
func Rotate(trans Mat4, theta float64, axis Vec3) Mat4 {
	id := Identity4()
	sin, cos := math.Sincos(theta)
	sin32, cos32 := float32(sin), float32(cos)
	cos_m_1 := 1 - cos32

	id[0][0] = cos32 + (axis[0]*axis[0])*cos_m_1
	id[1][0] = axis[1]*axis[0]*cos_m_1 + axis[2]*sin32
	id[2][0] = axis[2]*axis[0]*cos_m_1 - axis[1]*sin32

	id[0][1] = axis[0]*axis[1]*cos_m_1 - axis[2]*sin32
	id[1][1] = cos32 + (axis[1]*axis[1])*cos_m_1
	id[2][1] = axis[2]*axis[1]*cos_m_1 + axis[0]*sin32

	id[0][2] = axis[0]*axis[2]*cos_m_1 + axis[1]*sin32
	id[1][2] = axis[1]*axis[2]*cos_m_1 - axis[0]*sin32
	id[2][2] = cos32 + (axis[2]*axis[2])*cos_m_1

	return trans.Mul(id)
}
