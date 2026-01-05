package matrices

import . "scop/src/render/vectors"

type Mat3 [3][3]float32
type Mat3x2 [3][2]float32
type Mat3x4 [3][4]float32

func (m Mat3) ToFlat() [3 * 3]float32 {
	return [3 * 3]float32{
		m[0][0], m[0][1], m[0][2],
		m[1][0], m[1][1], m[1][2],
		m[2][0], m[2][1], m[2][2],
	}
}

func (m *Mat3x2) ToFlat() [3 * 2]float32 {
	return [3 * 2]float32{
		m[0][0], m[0][1],
		m[1][0], m[1][1],
		m[2][0], m[2][1],
	}
}

func (m Mat3x4) ToFlat() [3 * 4]float32 {
	return [3 * 4]float32{
		m[0][0], m[0][1], m[0][2], m[0][3],
		m[1][0], m[1][1], m[1][2], m[1][3],
		m[2][0], m[2][1], m[2][2], m[2][3],
	}
}

func (m Mat3) Add(rhs Mat3) Mat3 {
	return Mat3{
		{m[0][0] + rhs[0][0], m[0][1] + rhs[0][1], m[0][2] + rhs[0][2]},
		{m[1][0] + rhs[1][0], m[1][1] + rhs[1][1], m[1][2] + rhs[1][2]},
		{m[2][0] + rhs[2][0], m[2][1] + rhs[2][1], m[2][2] + rhs[2][2]},
	}
}

func (m Mat3x2) Add(rhs Mat3x2) Mat3x2 {
	return Mat3x2{
		{m[0][0] + rhs[0][0], m[0][1] + rhs[0][1]},
		{m[1][0] + rhs[1][0], m[1][1] + rhs[1][1]},
		{m[2][0] + rhs[2][0], m[2][1] + rhs[2][1]},
	}
}

func (m Mat3x4) Add(rhs Mat3x4) Mat3x4 {
	return Mat3x4{
		{m[0][0] + rhs[0][0], m[0][1] + rhs[0][1], m[0][2] + rhs[0][2], m[0][3] + rhs[0][3]},
		{m[1][0] + rhs[1][0], m[1][1] + rhs[1][1], m[1][2] + rhs[1][2], m[1][3] + rhs[1][3]},
		{m[2][0] + rhs[2][0], m[2][1] + rhs[2][1], m[2][2] + rhs[2][2], m[2][3] + rhs[2][3]},
	}
}

func (m Mat3) AddScalar(k float32) Mat3 {
	return Mat3{
		{m[0][0] + k, m[0][1] + k, m[0][2] + k},
		{m[1][0] + k, m[1][1] + k, m[1][2] + k},
		{m[2][0] + k, m[2][1] + k, m[2][2] + k},
	}
}

func (m Mat3x2) AddScalar(k float32) Mat3x2 {
	return Mat3x2{
		{m[0][0] + k, m[0][1] + k},
		{m[1][0] + k, m[1][1] + k},
		{m[2][0] + k, m[2][1] + k},
	}
}

func (m Mat3x4) AddScalar(k float32) Mat3x4 {
	return Mat3x4{
		{m[0][0] + k, m[0][1] + k, m[0][2] + k, m[0][3] + k},
		{m[1][0] + k, m[1][1] + k, m[1][2] + k, m[1][3] + k},
		{m[2][0] + k, m[2][1] + k, m[2][2] + k, m[2][3] + k},
	}
}

func (m Mat3) Sub(rhs Mat3) Mat3 {
	return Mat3{
		{m[0][0] - rhs[0][0], m[0][1] - rhs[0][1], m[0][2] - rhs[0][2]},
		{m[1][0] - rhs[1][0], m[1][1] - rhs[1][1], m[1][2] - rhs[1][2]},
		{m[2][0] - rhs[2][0], m[2][1] - rhs[2][1], m[2][2] - rhs[2][2]},
	}
}

func (m Mat3x2) Sub(rhs Mat3x2) Mat3x2 {
	return Mat3x2{
		{m[0][0] - rhs[0][0], m[0][1] - rhs[0][1]},
		{m[1][0] - rhs[1][0], m[1][1] - rhs[1][1]},
		{m[2][0] - rhs[2][0], m[2][1] - rhs[2][1]},
	}
}

func (m Mat3x4) Sub(rhs Mat3x4) Mat3x4 {
	return Mat3x4{
		{m[0][0] - rhs[0][0], m[0][1] - rhs[0][1], m[0][2] - rhs[0][2], m[0][3] - rhs[0][3]},
		{m[1][0] - rhs[1][0], m[1][1] - rhs[1][1], m[1][2] - rhs[1][2], m[1][3] - rhs[1][3]},
		{m[2][0] - rhs[2][0], m[2][1] - rhs[2][1], m[2][2] - rhs[2][2], m[2][3] - rhs[2][3]},
	}
}

func (m Mat3) SubScalar(k float32) Mat3 {
	return Mat3{
		{m[0][0] - k, m[0][1] - k, m[0][2] - k},
		{m[1][0] - k, m[1][1] - k, m[1][2] - k},
		{m[2][0] - k, m[2][1] - k, m[2][2] - k},
	}
}

func (m Mat3x2) SubScalar(k float32) Mat3x2 {
	return Mat3x2{
		{m[0][0] - k, m[0][1] - k},
		{m[1][0] - k, m[1][1] - k},
		{m[2][0] - k, m[2][1] - k},
	}
}

func (m Mat3x4) SubScalar(k float32) Mat3x4 {
	return Mat3x4{
		{m[0][0] - k, m[0][1] - k, m[0][2] - k, m[0][3] - k},
		{m[1][0] - k, m[1][1] - k, m[1][2] - k, m[1][3] - k},
		{m[2][0] - k, m[2][1] - k, m[2][2] - k, m[2][3] - k},
	}
}

func (m Mat3) MulScalar(k float32) Mat3 {
	return Mat3{
		{m[0][0] * k, m[0][1] * k, m[0][2] * k},
		{m[1][0] * k, m[1][1] * k, m[1][2] * k},
		{m[2][0] * k, m[2][1] * k, m[2][2] * k},
	}
}

func (m Mat3x2) MulScalar(k float32) Mat3x2 {
	return Mat3x2{
		{m[0][0] * k, m[0][1] * k},
		{m[1][0] * k, m[1][1] * k},
		{m[2][0] * k, m[2][1] * k},
	}
}

func (m Mat3x4) MulScalar(k float32) Mat3x4 {
	return Mat3x4{
		{m[0][0] * k, m[0][1] * k, m[0][2] * k, m[0][3] * k},
		{m[1][0] * k, m[1][1] * k, m[1][2] * k, m[1][3] * k},
		{m[2][0] * k, m[2][1] * k, m[2][2] * k, m[2][3] * k},
	}
}

func (m Mat3) Mul3x1(rhs Vec3) Vec3 {
	return Vec3{
		m[0][0]*rhs[0] + m[0][1]*rhs[1] + m[0][2]*rhs[2],
		m[1][0]*rhs[0] + m[1][1]*rhs[1] + m[1][2]*rhs[2],
		m[2][0]*rhs[0] + m[2][1]*rhs[1] + m[2][2]*rhs[2],
	}
}

func (m Mat3) Mul(rhs Mat3) Mat3 {
	return Mat3{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2] + m[0][2]*rhs[2][2]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2] + m[1][2]*rhs[2][2]},
		{m[2][0]*rhs[0][0] + m[2][1]*rhs[1][0] + m[2][2]*rhs[2][0], m[2][0]*rhs[0][1] + m[2][1]*rhs[1][1] + m[2][2]*rhs[2][1], m[2][0]*rhs[0][2] + m[2][1]*rhs[1][2] + m[2][2]*rhs[2][2]},
	}
}

func (m Mat3) Mul3x2(rhs Mat3x2) Mat3x2 {
	return Mat3x2{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1]},
		{m[2][0]*rhs[0][0] + m[2][1]*rhs[1][0] + m[2][2]*rhs[2][0], m[2][0]*rhs[0][1] + m[2][1]*rhs[1][1] + m[2][2]*rhs[2][1]},
	}
}

func (m Mat3) Mul3x4(rhs Mat3x4) Mat3x4 {
	return Mat3x4{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2] + m[0][2]*rhs[2][2], m[0][0]*rhs[0][3] + m[0][1]*rhs[1][3] + m[0][2]*rhs[2][3]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2] + m[1][2]*rhs[2][2], m[1][0]*rhs[0][3] + m[1][1]*rhs[1][3] + m[1][2]*rhs[2][3]},
		{m[2][0]*rhs[0][0] + m[2][1]*rhs[1][0] + m[2][2]*rhs[2][0], m[2][0]*rhs[0][1] + m[2][1]*rhs[1][1] + m[2][2]*rhs[2][1], m[2][0]*rhs[0][2] + m[2][1]*rhs[1][2] + m[2][2]*rhs[2][2], m[2][0]*rhs[0][3] + m[2][1]*rhs[1][3] + m[2][2]*rhs[2][3]},
	}
}

func (m Mat3x2) Mul2x1(rhs Vec2) Vec3 {
	return Vec3{
		m[0][0]*rhs[0] + m[0][1]*rhs[1],
		m[1][0]*rhs[0] + m[1][1]*rhs[1],
		m[2][0]*rhs[0] + m[2][1]*rhs[1],
	}
}

func (m Mat3x2) Mul2x2(rhs Mat2) Mat3x2 {
	return Mat3x2{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1]},
		{m[2][0]*rhs[0][0] + m[2][1]*rhs[1][0], m[2][0]*rhs[0][1] + m[2][1]*rhs[1][1]},
	}
}

func (m Mat3x2) Mul2x3(rhs Mat2x3) Mat3 {
	return Mat3{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2]},
		{m[2][0]*rhs[0][0] + m[2][1]*rhs[1][0], m[2][0]*rhs[0][1] + m[2][1]*rhs[1][1], m[2][0]*rhs[0][2] + m[2][1]*rhs[1][2]},
	}
}

func (m Mat3x2) Mul2x4(rhs Mat2x4) Mat3x4 {
	return Mat3x4{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2], m[0][0]*rhs[0][3] + m[0][1]*rhs[1][3]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2], m[1][0]*rhs[0][3] + m[1][1]*rhs[1][3]},
		{m[2][0]*rhs[0][0] + m[2][1]*rhs[1][0], m[2][0]*rhs[0][1] + m[2][1]*rhs[1][1], m[2][0]*rhs[0][2] + m[2][1]*rhs[1][2], m[2][0]*rhs[0][3] + m[2][1]*rhs[1][3]},
	}
}

func (m Mat3x4) Mul4x1(rhs Vec4) Vec3 {
	return Vec3{
		m[0][0]*rhs[0] + m[0][1]*rhs[1] + m[0][2]*rhs[2] + m[0][3]*rhs[3],
		m[1][0]*rhs[0] + m[1][1]*rhs[1] + m[1][2]*rhs[2] + m[1][3]*rhs[3],
		m[2][0]*rhs[0] + m[2][1]*rhs[1] + m[2][2]*rhs[2] + m[2][3]*rhs[3],
	}
}

func (m Mat3x4) Mul4x2(rhs Mat4x2) Mat3x2 {
	return Mat3x2{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0] + m[0][3]*rhs[3][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1] + m[0][3]*rhs[3][1]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0] + m[1][3]*rhs[3][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1] + m[1][3]*rhs[3][1]},
		{m[2][0]*rhs[0][0] + m[2][1]*rhs[1][0] + m[2][2]*rhs[2][0] + m[2][3]*rhs[3][0], m[2][0]*rhs[0][1] + m[2][1]*rhs[1][1] + m[2][2]*rhs[2][1] + m[2][3]*rhs[3][1]},
	}
}

func (m Mat3x4) Mul4x3(rhs Mat4x3) Mat3 {
	return Mat3{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0] + m[0][3]*rhs[3][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1] + m[0][3]*rhs[3][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2] + m[0][2]*rhs[2][2] + m[0][3]*rhs[3][2]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0] + m[1][3]*rhs[3][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1] + m[1][3]*rhs[3][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2] + m[1][2]*rhs[2][2] + m[1][3]*rhs[3][2]},
		{m[2][0]*rhs[0][0] + m[2][1]*rhs[1][0] + m[2][2]*rhs[2][0] + m[2][3]*rhs[3][0], m[2][0]*rhs[0][1] + m[2][1]*rhs[1][1] + m[2][2]*rhs[2][1] + m[2][3]*rhs[3][1], m[2][0]*rhs[0][2] + m[2][1]*rhs[1][2] + m[2][2]*rhs[2][2] + m[2][3]*rhs[3][2]},
	}
}

func (m Mat3x4) Mul4x4(rhs Mat4) Mat3x4 {
	return Mat3x4{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0] + m[0][3]*rhs[3][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1] + m[0][3]*rhs[3][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2] + m[0][2]*rhs[2][2] + m[0][3]*rhs[3][2], m[0][0]*rhs[0][3] + m[0][1]*rhs[1][3] + m[0][2]*rhs[2][3] + m[0][3]*rhs[3][3]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0] + m[1][3]*rhs[3][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1] + m[1][3]*rhs[3][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2] + m[1][2]*rhs[2][2] + m[1][3]*rhs[3][2], m[1][0]*rhs[0][3] + m[1][1]*rhs[1][3] + m[1][2]*rhs[2][3] + m[1][3]*rhs[3][3]},
		{m[2][0]*rhs[0][0] + m[2][1]*rhs[1][0] + m[2][2]*rhs[2][0] + m[2][3]*rhs[3][0], m[2][0]*rhs[0][1] + m[2][1]*rhs[1][1] + m[2][2]*rhs[2][1] + m[2][3]*rhs[3][1], m[2][0]*rhs[0][2] + m[2][1]*rhs[1][2] + m[2][2]*rhs[2][2] + m[2][3]*rhs[3][2], m[2][0]*rhs[0][3] + m[2][1]*rhs[1][3] + m[2][2]*rhs[2][3] + m[2][3]*rhs[3][3]},
	}
}

func Identity3() Mat3 {
	return Mat3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

func (m Mat3) ToMat2() Mat2 {
	return Mat2{
		{m[0][0], m[0][1]},
		{m[1][0], m[1][1]},
	}
}

func (m Mat3) ToMat4(w float32) Mat4 {
	return Mat4{
		{m[0][0], m[0][1], m[0][2], 0},
		{m[1][0], m[1][1], m[1][2], 0},
		{m[2][0], m[2][1], m[2][2], 0},
		{0, 0, 0, w},
	}
}
