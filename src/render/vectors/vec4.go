package vectors

import "math"

type Vec4 [4]float32

func (v Vec4) Add(rhs Vec4) Vec4 {
	return Vec4{v[0] + rhs[0], v[1] + rhs[1], v[2] + rhs[2], v[3] + rhs[3]}
}

func (v Vec4) AddScalar(k float32) Vec4 {
	return Vec4{v[0] + k, v[1] + k, v[2] + k, v[3] + k}
}

func (v Vec4) Sub(rhs Vec4) Vec4 {
	return Vec4{v[0] - rhs[0], v[1] - rhs[1], v[2] - rhs[2], v[3] - rhs[3]}
}

func (v Vec4) SubScalar(k float32) Vec4 {
	return Vec4{v[0] - k, v[1] - k, v[2] - k, v[3] - k}
}

func (v Vec4) Mul(rhs Vec4) Vec4 {
	return Vec4{v[0] * rhs[0], v[1] * rhs[1], v[2] * rhs[2], v[3] * rhs[3]}
}

func (v Vec4) MulScalar(k float32) Vec4 {
	return Vec4{v[0] * k, v[1] * k, v[2] * k, v[3] * k}
}

func (v Vec4) Div(rhs Vec4) Vec4 {
	if rhs[0] == 0 || rhs[1] == 0 || rhs[2] == 0 || rhs[3] == 0 {
		panic("division by zero")
	}
	return Vec4{v[0] / rhs[0], v[1] / rhs[1], v[2] / rhs[2], v[3] / rhs[3]}
}

func (v Vec4) DivScalar(k float32) Vec4 {
	if k == 0 {
		panic("division by zero")
	}
	return Vec4{v[0] / k, v[1] / k, v[2] / k, v[3] / k}
}

func (v Vec4) Neg() Vec4 {
	return Vec4{-v[0], -v[1], -v[2], -v[3]}
}

func (v Vec4) Magnitude() float32 {
	return float32(math.Sqrt(
		float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3]),
	))
}

func (v Vec4) DotProduct(rhs Vec4) float32 {
	return v[0]*rhs[0] + v[1]*rhs[1] + v[2]*rhs[2] + v[3]*rhs[3]
}

func (v Vec4) ToVec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec4) ToVec3() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}
