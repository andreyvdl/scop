package transformations

import (
	"math"
	. "scop/src/render/matrices"
	. "scop/src/render/vectors"
)

func RotateX(theta float64, v Vec4) Vec4 {
	r := Identity4()

	r[5] = float32(math.Cos(theta))
	r[6] = float32(-math.Sin(theta))
	r[9] = float32(math.Sin(theta))
	r[10] = float32(math.Cos(theta))
	return r.Mul4x1(v)
}

func RotateY(theta float64, v Vec4) Vec4 {
	r := Identity4()

	r[0] = float32(math.Cos(theta))
	r[2] = float32(math.Sin(theta))
	r[8] = float32(-math.Sin(theta))
	r[10] = float32(math.Cos(theta))
	return r.Mul4x1(v)
}

func RotateZ(theta float64, v Vec4) Vec4 {
	r := Identity4()

	r[0] = float32(math.Cos(theta))
	r[1] = float32(-math.Sin(theta))
	r[4] = float32(math.Sin(theta))
	r[5] = float32(math.Cos(theta))
	return r.Mul4x1(v)
}
