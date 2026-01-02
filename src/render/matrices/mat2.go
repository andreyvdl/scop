package matrices

import . "scop/src/render/vectors"

type Mat2 [2 * 2]float32
type Mat2x3 [2 * 3]float32
type Mat2x4 [2 * 4]float32

func (m Mat2) Add(rhs Mat2) Mat2 {
	return Mat2{
		m[0] + rhs[0], m[1] + rhs[1],
		m[2] + rhs[2], m[3] + rhs[3],
	}
}

func (m Mat2x3) Add(rhs Mat2x3) Mat2x3 {
	return Mat2x3{
		m[0] + rhs[0], m[1] + rhs[1], m[2] + rhs[2],
		m[3] + rhs[3], m[4] + rhs[4], m[5] + rhs[5],
	}
}

func (m Mat2x4) Add(rhs Mat2x4) Mat2x4 {
	return Mat2x4{
		m[0] + rhs[0], m[1] + rhs[1], m[2] + rhs[2], m[3] + rhs[3],
		m[4] + rhs[4], m[5] + rhs[5], m[6] + rhs[6], m[7] + rhs[7],
	}
}

func (m Mat2) AddScalar(k float32) Mat2 {
	return Mat2{
		m[0] + k, m[1] + k,
		m[2] + k, m[3] + k,
	}
}

func (m Mat2x3) AddScalar(k float32) Mat2x3 {
	return Mat2x3{
		m[0] + k, m[1] + k, m[2] + k,
		m[3] + k, m[4] + k, m[5] + k,
	}
}

func (m Mat2x4) AddScalar(k float32) Mat2x4 {
	return Mat2x4{
		m[0] + k, m[1] + k, m[2] + k, m[3] + k,
		m[4] + k, m[5] + k, m[6] + k, m[7] + k,
	}
}

func (m Mat2) Sub(rhs Mat2) Mat2 {
	return Mat2{
		m[0] - rhs[0], m[1] - rhs[1],
		m[2] - rhs[2], m[3] - rhs[3],
	}
}

func (m Mat2x3) Sub(rhs Mat2x3) Mat2x3 {
	return Mat2x3{
		m[0] - rhs[0], m[1] - rhs[1], m[2] - rhs[2],
		m[3] - rhs[3], m[4] - rhs[4], m[5] - rhs[5],
	}
}

func (m Mat2x4) Sub(rhs Mat2x4) Mat2x4 {
	return Mat2x4{
		m[0] - rhs[0], m[1] - rhs[1], m[2] - rhs[2], m[3] - rhs[3],
		m[4] - rhs[4], m[5] - rhs[5], m[6] - rhs[6], m[7] - rhs[7],
	}
}

func (m Mat2) SubScalar(k float32) Mat2 {
	return Mat2{
		m[0] - k, m[1] - k,
		m[2] - k, m[3] - k,
	}
}

func (m Mat2x3) SubScalar(k float32) Mat2x3 {
	return Mat2x3{
		m[0] - k, m[1] - k, m[2] - k,
		m[3] - k, m[4] - k, m[5] - k,
	}
}

func (m Mat2x4) SubScalar(k float32) Mat2x4 {
	return Mat2x4{
		m[0] - k, m[1] - k, m[2] - k, m[3] - k,
		m[4] - k, m[5] - k, m[6] - k, m[7] - k,
	}
}

func (m Mat2) MulScalar(k float32) Mat2 {
	return Mat2{
		m[0] * k, m[1] * k,
		m[2] * k, m[3] * k,
	}
}

func (m Mat2x3) MulScalar(k float32) Mat2x3 {
	return Mat2x3{
		m[0] * k, m[1] * k, m[2] * k,
		m[3] * k, m[4] * k, m[5] * k,
	}
}

func (m Mat2x4) MulScalar(k float32) Mat2x4 {
	return Mat2x4{
		m[0] * k, m[1] * k, m[2] * k, m[3] * k,
		m[4] * k, m[5] * k, m[6] * k, m[7] * k,
	}
}

func (m Mat2) Mul2x1(rhs Vec2) Vec2 {
	return Vec2{
		m[0]*rhs[0] + m[1]*rhs[1],
		m[2]*rhs[0] + m[3]*rhs[1],
	}
}

func (m Mat2) Mul(rhs Mat2) Mat2 {
	return Mat2{
		m[0]*rhs[0] + m[1]*rhs[2], m[0]*rhs[1] + m[1]*rhs[3],
		m[2]*rhs[0] + m[3]*rhs[2], m[2]*rhs[1] + m[3]*rhs[3],
	}
}

func (m Mat2) Mul2x3(rhs Mat2x3) Mat2x3 {
	return Mat2x3{
		m[0]*rhs[0] + m[1]*rhs[3], m[0]*rhs[1] + m[1]*rhs[4], m[0]*rhs[2] + m[1]*rhs[5],
		m[2]*rhs[0] + m[3]*rhs[3], m[2]*rhs[1] + m[3]*rhs[4], m[2]*rhs[2] + m[3]*rhs[5],
	}
}

func (m Mat2) Mul2x4(rhs Mat2x4) Mat2x4 {
	return Mat2x4{
		m[0]*rhs[0] + m[1]*rhs[4], m[0]*rhs[1] + m[1]*rhs[5], m[0]*rhs[2] + m[1]*rhs[6], m[0]*rhs[3] + m[1]*rhs[7],
		m[2]*rhs[0] + m[3]*rhs[4], m[2]*rhs[1] + m[3]*rhs[5], m[2]*rhs[2] + m[3]*rhs[6], m[2]*rhs[3] + m[3]*rhs[7],
	}
}

func Identity2() Mat2 {
	return Mat2{
		1, 0,
		0, 1,
	}
}

func (m Mat2) ToMat3(z float32) Mat3 {
	return Mat3{
		m[0], m[1], 0,
		m[2], m[3], 0,
		0, 0, z,
	}
}

func (m Mat2) ToMat4(z, w float32) Mat4 {
	return Mat4{
		m[0], m[1], 0, 0,
		m[2], m[3], 0, 0,
		0, 0, z, 0,
		0, 0, 0, w,
	}
}
