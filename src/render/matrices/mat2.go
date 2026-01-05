package matrices

import . "scop/src/render/vectors"

type Mat2 [2][2]float32
type Mat2x3 [2][3]float32
type Mat2x4 [2][4]float32

func (m Mat2) ToFlat() [2 * 2]float32 {
	return [2 * 2]float32{
		m[0][0], m[0][1],
		m[1][0], m[1][1],
	}
}

func (m Mat2x3) ToFlat() [2 * 3]float32 {
	return [2 * 3]float32{
		m[0][0], m[0][1], m[0][2],
		m[1][0], m[1][1], m[1][2],
	}
}

func (m Mat2x4) ToFlat() [2 * 4]float32 {
	return [2 * 4]float32{
		m[0][0], m[0][1], m[0][2], m[0][3],
		m[1][0], m[1][1], m[1][2], m[1][3],
	}
}

func (m Mat2) Add(rhs Mat2) Mat2 {
	return Mat2{
		{m[0][0] + rhs[0][0], m[0][1] + rhs[0][1]},
		{m[1][0] + rhs[1][0], m[1][1] + rhs[1][1]},
	}
}

func (m Mat2x3) Add(rhs Mat2x3) Mat2x3 {
	return Mat2x3{
		{m[0][0] + rhs[0][0], m[0][1] + rhs[0][1], m[0][2] + rhs[0][2]},
		{m[1][0] + rhs[1][0], m[1][1] + rhs[1][1], m[1][2] + rhs[1][2]},
	}
}

func (m Mat2x4) Add(rhs Mat2x4) Mat2x4 {
	return Mat2x4{
		{m[0][0] + rhs[0][0], m[0][1] + rhs[0][1], m[0][2] + rhs[0][2], m[0][3] + rhs[0][3]},
		{m[1][0] + rhs[1][0], m[1][1] + rhs[1][1], m[1][2] + rhs[1][2], m[1][3] + rhs[1][3]},
	}
}

func (m Mat2) AddScalar(k float32) Mat2 {
	return Mat2{
		{m[0][0] + k, m[0][1] + k},
		{m[1][0] + k, m[1][1] + k},
	}
}

func (m Mat2x3) AddScalar(k float32) Mat2x3 {
	return Mat2x3{
		{m[0][0] + k, m[0][1] + k, m[0][2] + k},
		{m[1][0] + k, m[1][1] + k, m[1][2] + k},
	}
}

func (m Mat2x4) AddScalar(k float32) Mat2x4 {
	return Mat2x4{
		{m[0][0] + k, m[0][1] + k, m[0][2] + k, m[0][3] + k},
		{m[1][0] + k, m[1][1] + k, m[1][2] + k, m[1][3] + k},
	}
}

func (m Mat2) Sub(rhs Mat2) Mat2 {
	return Mat2{
		{m[0][0] - rhs[0][0], m[0][1] - rhs[0][1]},
		{m[1][0] - rhs[1][0], m[1][1] - rhs[1][1]},
	}
}

func (m Mat2x3) Sub(rhs Mat2x3) Mat2x3 {
	return Mat2x3{
		{m[0][0] - rhs[0][0], m[0][1] - rhs[0][1], m[0][2] - rhs[0][2]},
		{m[1][0] - rhs[1][0], m[1][1] - rhs[1][1], m[1][2] - rhs[1][2]},
	}
}

func (m Mat2x4) Sub(rhs Mat2x4) Mat2x4 {
	return Mat2x4{
		{m[0][0] - rhs[0][0], m[0][1] - rhs[0][1], m[0][2] - rhs[0][2], m[0][3] - rhs[0][3]},
		{m[1][0] - rhs[1][0], m[1][1] - rhs[1][1], m[1][2] - rhs[1][2], m[1][3] - rhs[1][3]},
	}
}

func (m Mat2) SubScalar(k float32) Mat2 {
	return Mat2{
		{m[0][0] - k, m[0][1] - k},
		{m[1][0] - k, m[1][1] - k},
	}
}

func (m Mat2x3) SubScalar(k float32) Mat2x3 {
	return Mat2x3{
		{m[0][0] - k, m[0][1] - k, m[0][2] - k},
		{m[1][0] - k, m[1][1] - k, m[1][2] - k},
	}
}

func (m Mat2x4) SubScalar(k float32) Mat2x4 {
	return Mat2x4{
		{m[0][0] - k, m[0][1] - k, m[0][2] - k, m[0][3] - k},
		{m[1][0] - k, m[1][1] - k, m[1][2] - k, m[1][3] - k},
	}
}

func (m Mat2) MulScalar(k float32) Mat2 {
	return Mat2{
		{m[0][0] * k, m[0][1] * k},
		{m[1][0] * k, m[1][1] * k},
	}
}

func (m Mat2x3) MulScalar(k float32) Mat2x3 {
	return Mat2x3{
		{m[0][0] * k, m[0][1] * k, m[0][2] * k},
		{m[1][0] * k, m[1][1] * k, m[1][2] * k},
	}
}

func (m Mat2x4) MulScalar(k float32) Mat2x4 {
	return Mat2x4{
		{m[0][0] * k, m[0][1] * k, m[0][2] * k, m[0][3] * k},
		{m[1][0] * k, m[1][1] * k, m[1][2] * k, m[1][3] * k},
	}
}

func (m Mat2) Mul2x1(rhs Vec2) Vec2 {
	return Vec2{
		m[0][0]*rhs[0] + m[0][1]*rhs[1],
		m[1][0]*rhs[0] + m[1][1]*rhs[1],
	}
}

func (m Mat2) Mul(rhs Mat2) Mat2 {
	return Mat2{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1]},
	}
}

func (m Mat2) Mul2x3(rhs Mat2x3) Mat2x3 {
	return Mat2x3{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2]},
	}
}

func (m Mat2) Mul2x4(rhs Mat2x4) Mat2x4 {
	return Mat2x4{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2], m[0][0]*rhs[0][3] + m[0][1]*rhs[1][3]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2], m[1][0]*rhs[0][3] + m[1][1]*rhs[1][3]},
	}
}

func (m Mat2x3) Mul3x1(rhs Vec3) Vec3 {
	return Vec3{
		m[0][0]*rhs[0] + m[0][1]*rhs[1] + m[0][2]*rhs[2],
		m[1][0]*rhs[0] + m[1][1]*rhs[1] + m[1][2]*rhs[2],
	}
}

func (m Mat2x3) Mul3x2(rhs Mat3x2) Mat3 {
	return Mat3{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1]},
	}
}

func (m Mat2x3) Mul3x3(rhs Mat3) Mat2x3 {
	return Mat2x3{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2] + m[0][2]*rhs[2][2]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2] + m[1][2]*rhs[2][2]},
	}
}

func (m Mat2x3) Mul3x4(rhs Mat3x4) Mat2x4 {
	return Mat2x4{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2] + m[0][2]*rhs[2][2], m[0][0]*rhs[0][3] + m[0][1]*rhs[1][3] + m[0][2]*rhs[2][3]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2] + m[1][2]*rhs[2][2], m[1][0]*rhs[0][3] + m[1][1]*rhs[1][3] + m[1][2]*rhs[2][3]},
	}
}

func (m Mat2x4) Mul4x1(rhs Vec4) Vec2 {
	return Vec2{
		m[0][0]*rhs[0] + m[0][1]*rhs[1] + m[0][2]*rhs[2] + m[0][3]*rhs[3],
		m[1][0]*rhs[0] + m[1][1]*rhs[1] + m[1][2]*rhs[2] + m[1][3]*rhs[3],
	}
}

func (m Mat2x4) Mul4x2(rhs Mat4x2) Mat2 {
	return Mat2{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0] + m[0][3]*rhs[3][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1] + m[0][3]*rhs[3][1]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0] + m[1][3]*rhs[3][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1] + m[1][3]*rhs[3][1]},
	}
}

func (m Mat2x4) Mul4x3(rhs Mat4x3) Mat2x3 {
	return Mat2x3{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0] + m[0][3]*rhs[3][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1] + m[0][3]*rhs[3][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2] + m[0][2]*rhs[2][2] + m[0][3]*rhs[3][2]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0] + m[1][3]*rhs[3][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1] + m[1][3]*rhs[3][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2] + m[1][2]*rhs[2][2] + m[0][3]*rhs[3][2]},
	}
}

func (m Mat2x4) Mul4x4(rhs Mat4) Mat2x4 {
	return Mat2x4{
		{m[0][0]*rhs[0][0] + m[0][1]*rhs[1][0] + m[0][2]*rhs[2][0] + m[0][3]*rhs[3][0], m[0][0]*rhs[0][1] + m[0][1]*rhs[1][1] + m[0][2]*rhs[2][1] + m[0][3]*rhs[3][1], m[0][0]*rhs[0][2] + m[0][1]*rhs[1][2] + m[0][2]*rhs[2][2] + m[0][3]*rhs[3][2], m[0][0]*rhs[0][3] + m[0][1]*rhs[1][3] + m[0][2]*rhs[2][3] + m[0][3]*rhs[3][3]},
		{m[1][0]*rhs[0][0] + m[1][1]*rhs[1][0] + m[1][2]*rhs[2][0] + m[1][3]*rhs[3][0], m[1][0]*rhs[0][1] + m[1][1]*rhs[1][1] + m[1][2]*rhs[2][1] + m[1][3]*rhs[3][1], m[1][0]*rhs[0][2] + m[1][1]*rhs[1][2] + m[1][2]*rhs[2][2] + m[1][3]*rhs[3][2], m[1][0]*rhs[0][3] + m[1][1]*rhs[1][3] + m[1][2]*rhs[2][3] + m[1][3]*rhs[3][3]},
	}
}

func Identity2() Mat2 {
	return Mat2{
		{1, 0},
		{0, 1},
	}
}

func (m Mat2) ToMat3(z float32) Mat3 {
	return Mat3{
		{m[0][0], m[0][1], 0},
		{m[1][0], m[1][1], 0},
		{0, 0, z},
	}
}

func (m Mat2) ToMat4(z, w float32) Mat4 {
	return Mat4{
		{m[0][0], m[0][1], 0, 0},
		{m[1][0], m[1][1], 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, w},
	}
}
