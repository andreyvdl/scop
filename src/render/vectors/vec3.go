package vectors

import "math"

type Vec3 [3]float32

func (v Vec3) Add(rhs Vec3) Vec3 {
	return Vec3{v[0] + rhs[0], v[1] + rhs[1], v[2] + rhs[2]}
}

func (v Vec3) AddScalar(k float32) Vec3 {
	return Vec3{v[0] + k, v[1] + k, v[2] + k}
}

func (v Vec3) Sub(rhs Vec3) Vec3 {
	return Vec3{v[0] - rhs[0], v[1] - rhs[1], v[2] - rhs[2]}
}

func (v Vec3) SubScalar(k float32) Vec3 {
	return Vec3{v[0] - k, v[1] - k, v[2] - k}
}

func (v Vec3) Mul(rhs Vec3) Vec3 {
	return Vec3{v[0] * rhs[0], v[1] * rhs[1], v[2] * rhs[2]}
}

func (v Vec3) MulScalar(k float32) Vec3 {
	return Vec3{v[0] * k, v[1] * k, v[2] * k}
}

func (v Vec3) Div(rhs Vec3) Vec3 {
	if rhs[0] == 0 || rhs[1] == 0 || rhs[2] == 0 {
		panic("division by zero")
	}
	return Vec3{v[0] / rhs[0], v[1] / rhs[1], v[2] / rhs[2]}
}

func (v Vec3) DivScalar(k float32) Vec3 {
	if k == 0 {
		panic("division by zero")
	}
	return Vec3{v[0] / k, v[1] / k, v[2] / k}
}

func (v Vec3) Neg() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

func (v Vec3) Magnitude() float32 {
	return float32(math.Sqrt(
		float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2]),
	))
}

func (v Vec3) DotProduct(rhs Vec3) float32 {
	return v[0]*rhs[0] + v[1]*rhs[1] + v[2]*rhs[2]
}

func (v Vec3) CrossProduct(rhs Vec3) Vec3 {
	return Vec3{
		v[1]*rhs[2] - v[2]*rhs[1],
		v[2]*rhs[0] - v[0]*rhs[2],
		v[0]*rhs[1] - v[1]*rhs[0],
	}
}

func (v Vec3) ToVec4(w float32) Vec4 {
	return Vec4{v[0], v[1], v[2], w}
}

func (v Vec3) ToVec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec3) Normalize() Vec3 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vec3{0, 0, 0}
	}
	return v.DivScalar(mag)
}
