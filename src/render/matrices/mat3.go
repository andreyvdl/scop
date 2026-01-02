package matrices

import . "scop/src/render/vectors"

type Mat3 [3 * 3]float32
type Mat3x2 [3 * 2]float32
type Mat3x4 [3 * 4]float32

func (m Mat3) Add(rhs Mat3) Mat3 {
	return Mat3{
		m[0] + rhs[0], m[1] + rhs[1], m[2] + rhs[2],
		m[3] + rhs[3], m[4] + rhs[4], m[5] + rhs[5],
		m[6] + rhs[6], m[7] + rhs[7], m[8] + rhs[8],
	}
}

func (m Mat3x2) Add(rhs Mat3x2) Mat3x2 {
	return Mat3x2{
		m[0] + rhs[0], m[1] + rhs[1],
		m[2] + rhs[2], m[3] + rhs[3],
		m[4] + rhs[4], m[5] + rhs[5],
	}
}

func (m Mat3x4) Add(rhs Mat3x4) Mat3x4 {
	return Mat3x4{
		m[0] + rhs[0], m[1] + rhs[1], m[2] + rhs[2], m[3] + rhs[3],
		m[4] + rhs[4], m[5] + rhs[5], m[6] + rhs[6], m[7] + rhs[7],
		m[8] + rhs[8], m[9] + rhs[9], m[10] + rhs[10], m[11] + rhs[11],
	}
}

func (m Mat3) AddScalar(k float32) Mat3 {
	return Mat3{
		m[0] + k, m[1] + k, m[2] + k,
		m[3] + k, m[4] + k, m[5] + k,
		m[6] + k, m[7] + k, m[8] + k,
	}
}

func (m Mat3x2) AddScalar(k float32) Mat3x2 {
	return Mat3x2{
		m[0] + k, m[1] + k,
		m[2] + k, m[3] + k,
		m[4] + k, m[5] + k,
	}
}

func (m Mat3x4) AddScalar(k float32) Mat3x4 {
	return Mat3x4{
		m[0] + k, m[1] + k, m[2] + k, m[3] + k,
		m[4] + k, m[5] + k, m[6] + k, m[7] + k,
		m[8] + k, m[9] + k, m[10] + k, m[11] + k,
	}
}

func (m Mat3) Sub(rhs Mat3) Mat3 {
	return Mat3{
		m[0] - rhs[0], m[1] - rhs[1], m[2] - rhs[2],
		m[3] - rhs[3], m[4] - rhs[4], m[5] - rhs[5],
		m[6] - rhs[6], m[7] - rhs[7], m[8] - rhs[8],
	}
}

func (m Mat3x2) Sub(rhs Mat3x2) Mat3x2 {
	return Mat3x2{
		m[0] - rhs[0], m[1] - rhs[1],
		m[2] - rhs[2], m[3] - rhs[3],
		m[4] - rhs[4], m[5] - rhs[5],
	}
}

func (m Mat3x4) Sub(rhs Mat3x4) Mat3x4 {
	return Mat3x4{
		m[0] - rhs[0], m[1] - rhs[1], m[2] - rhs[2], m[3] - rhs[3],
		m[4] - rhs[4], m[5] - rhs[5], m[6] - rhs[6], m[7] - rhs[7],
		m[8] - rhs[8], m[9] - rhs[9], m[10] - rhs[10], m[11] - rhs[11],
	}
}

func (m Mat3) SubScalar(k float32) Mat3 {
	return Mat3{
		m[0] - k, m[1] - k, m[2] - k,
		m[3] - k, m[4] - k, m[5] - k,
		m[6] - k, m[7] - k, m[8] - k,
	}
}

func (m Mat3x2) SubScalar(k float32) Mat3x2 {
	return Mat3x2{
		m[0] - k, m[1] - k,
		m[2] - k, m[3] - k,
		m[4] - k, m[5] - k,
	}
}

func (m Mat3x4) SubScalar(k float32) Mat3x4 {
	return Mat3x4{
		m[0] - k, m[1] - k, m[2] - k, m[3] - k,
		m[4] - k, m[5] - k, m[6] - k, m[7] - k,
		m[8] - k, m[9] - k, m[10] - k, m[11] - k,
	}
}

func (m Mat3) MulScalar(k float32) Mat3 {
	return Mat3{
		m[0] * k, m[1] * k, m[2] * k,
		m[3] * k, m[4] * k, m[5] * k,
		m[6] * k, m[7] * k, m[8] * k,
	}
}

func (m Mat3x2) MulScalar(k float32) Mat3x2 {
	return Mat3x2{
		m[0] * k, m[1] * k,
		m[2] * k, m[3] * k,
		m[4] * k, m[5] * k,
	}
}

func (m Mat3x4) MulScalar(k float32) Mat3x4 {
	return Mat3x4{
		m[0] * k, m[1] * k, m[2] * k, m[3] * k,
		m[4] * k, m[5] * k, m[6] * k, m[7] * k,
		m[8] * k, m[9] * k, m[10] * k, m[11] * k,
	}
}

func (m Mat3) Mul3x1(rhs Vec3) Vec3 {
	return Vec3{
		m[0]*rhs[0] + m[1]*rhs[1] + m[2]*rhs[2],
		m[3]*rhs[0] + m[4]*rhs[1] + m[5]*rhs[2],
		m[6]*rhs[0] + m[7]*rhs[1] + m[8]*rhs[2],
	}
}

func (m Mat3) Mul(rhs Mat3) Mat3 {
	return Mat3{
		m[0]*rhs[0] + m[1]*rhs[3] + m[2]*rhs[6], m[0]*rhs[1] + m[1]*rhs[4] + m[2]*rhs[7], m[0]*rhs[2] + m[1]*rhs[5] + m[2]*rhs[8],
		m[3]*rhs[0] + m[4]*rhs[3] + m[5]*rhs[6], m[3]*rhs[1] + m[4]*rhs[4] + m[5]*rhs[7], m[3]*rhs[2] + m[4]*rhs[5] + m[5]*rhs[8],
		m[6]*rhs[0] + m[7]*rhs[3] + m[8]*rhs[6], m[6]*rhs[1] + m[7]*rhs[4] + m[8]*rhs[7], m[6]*rhs[2] + m[7]*rhs[5] + m[8]*rhs[8],
	}
}

func (m Mat3) Mul3x2(rhs Mat3x2) Mat3x2 {
	return Mat3x2{
		m[0]*rhs[0] + m[1]*rhs[2] + m[2]*rhs[4], m[0]*rhs[1] + m[1]*rhs[3] + m[2]*rhs[5],
		m[3]*rhs[0] + m[4]*rhs[2] + m[5]*rhs[4], m[3]*rhs[1] + m[4]*rhs[3] + m[5]*rhs[5],
		m[6]*rhs[0] + m[7]*rhs[2] + m[8]*rhs[4], m[6]*rhs[1] + m[7]*rhs[3] + m[8]*rhs[5],
	}
}

func (m Mat3) Mul3x4(rhs Mat3x4) Mat3x4 {
	return Mat3x4{
		m[0]*rhs[0] + m[1]*rhs[4] + m[2]*rhs[8], m[0]*rhs[1] + m[1]*rhs[5] + m[2]*rhs[9], m[0]*rhs[2] + m[1]*rhs[6] + m[2]*rhs[10], m[0]*rhs[3] + m[1]*rhs[7] + m[2]*rhs[11],
		m[3]*rhs[0] + m[4]*rhs[4] + m[5]*rhs[8], m[3]*rhs[1] + m[4]*rhs[5] + m[5]*rhs[9], m[3]*rhs[2] + m[4]*rhs[6] + m[5]*rhs[10], m[3]*rhs[3] + m[4]*rhs[7] + m[5]*rhs[11],
		m[6]*rhs[0] + m[7]*rhs[4] + m[8]*rhs[8], m[6]*rhs[1] + m[7]*rhs[5] + m[8]*rhs[9], m[6]*rhs[2] + m[7]*rhs[6] + m[8]*rhs[10], m[6]*rhs[3] + m[7]*rhs[7] + m[8]*rhs[11],
	}
}

func Identity3() Mat3 {
	return Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func (m Mat3) ToMat2() Mat2 {
	return Mat2{
		m[0], m[1],
		m[3], m[4],
	}
}

func (m Mat3) ToMat4(w float32) Mat4 {
	return Mat4{
		m[0], m[1], m[2], 0,
		m[3], m[4], m[5], 0,
		m[6], m[7], m[8], 0,
		0, 0, 0, w,
	}
}
