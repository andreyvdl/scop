package vectors

import (
	"math"
)

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

func (v Vec4) Normalize() Vec4 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vec4{0, 0, 0, 0}
	}
	return v.DivScalar(mag)
}

type Quaternion Vec4

func (q Quaternion) Normalize() Quaternion {
	return Quaternion(Vec4(q).Normalize())
}

func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{-q[0], -q[1], -q[2], q[3]}
}

// WTF IS THIS SH???? JUST CALL CONJUGATE MF
func (q Quaternion) Inverse() Quaternion {
	return q.Conjugate()
}

func (q Quaternion) Mul(rhs Quaternion) Quaternion {
	return Quaternion{
		q[3]*rhs[0] + q[0]*rhs[3] + q[1]*rhs[2] - q[2]*rhs[1],
		q[3]*rhs[1] - q[0]*rhs[2] + q[1]*rhs[3] + q[2]*rhs[0],
		q[3]*rhs[2] + q[0]*rhs[1] - q[1]*rhs[0] + q[2]*rhs[3],
		q[3]*rhs[3] - q[0]*rhs[0] - q[1]*rhs[1] - q[2]*rhs[2],
	}
}

func (q Quaternion) RotateVec(v Vec3) Vec3 {
	q = q.Normalize()
	qv := Quaternion(v.ToVec4(0))
	return Vec4(q.Mul(qv).Mul(q.Conjugate())).ToVec3()
}

func FromAxisAngle(axis Vec3, angle float32) Quaternion {
	axis = axis.Normalize()

	sin, cos := math.Sincos(float64(angle / 2))
	sin32, cos32 := float32(sin), float32(cos)

	q := Quaternion{
		sin32 * axis[0],
		sin32 * axis[1],
		sin32 * axis[2],
		cos32,
	}
	return q
}

/*
This FUCKING function is hell
https://youtu.be/cIFtz7bPuUk?si=7Y85Su-vPq0Xkz6m
*/
func FromEuler(pitch, yaw, roll float32) Quaternion {
	rad := math.Pi / 180
	r_rad := float64(roll) * rad * 0.5
	p_rad := float64(pitch) * rad * 0.5
	y_rad := float64(yaw) * rad * 0.5

	rs, rc := math.Sincos(r_rad)
	ps, pc := math.Sincos(p_rad)
	ys, yc := math.Sincos(y_rad)

	return Quaternion{
		float32(ps*rc*yc - rs*ys*pc),
		float32(ps*rs*yc + ys*pc*rc),
		float32(ps*ys*rc + rs*pc*yc),
		float32(-ps*rs*ys + pc*rc*yc),
	}.Normalize()
}
