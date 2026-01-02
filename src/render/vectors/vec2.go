package vectors

import "math"

type Vec2 [2]float32

func (v Vec2) Add(rhs Vec2) Vec2 {
	return Vec2{v[0] + rhs[0], v[1] + rhs[1]}
}

func (v Vec2) AddScalar(k float32) Vec2 {
	return Vec2{v[0] + k, v[1] + k}
}

func (v Vec2) Sub(rhs Vec2) Vec2 {
	return Vec2{v[0] - rhs[0], v[1] - rhs[1]}
}

func (v Vec2) SubScalar(k float32) Vec2 {
	return Vec2{v[0] - k, v[1] - k}
}

func (v Vec2) Mul(rhs Vec2) Vec2 {
	return Vec2{v[0] * rhs[0], v[1] * rhs[1]}
}

func (v Vec2) MulScalar(k float32) Vec2 {
	return Vec2{v[0] * k, v[1] * k}
}

func (v Vec2) Div(rhs Vec2) Vec2 {
	if rhs[0] == 0 || rhs[1] == 0 {
		panic("division by zero")
	}
	return Vec2{v[0] / rhs[0], v[1] / rhs[1]}
}

func (v Vec2) DivScalar(k float32) Vec2 {
	if k == 0 {
		panic("division by zero")
	}
	return Vec2{v[0] / k, v[1] / k}
}

func (v Vec2) Neg() Vec2 {
	return Vec2{-v[0], -v[1]}
}

func (v Vec2) Magnitude() float32 {
	return float32(math.Sqrt(
		float64(v[0]*v[0] + v[1]*v[1]),
	))
}

func (v Vec2) DotProduct(rhs Vec2) float32 {
	return v[0]*rhs[0] + v[1]*rhs[1]
}

func (v Vec2) ToVec3(z float32) Vec3 {
	return Vec3{v[0], v[1], z}
}

func (v Vec2) ToVec4(z, w float32) Vec4 {
	return Vec4{v[0], v[1], z, w}
}
