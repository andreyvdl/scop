package matrices

import . "scop/src/render/vectors"

type Mat4 [4 * 4]float32
type Mat4x2 [4 * 2]float32
type Mat4x3 [4 * 3]float32

func (m Mat4) Add(rhs Mat4) Mat4 {
	return Mat4{
		m[0] + rhs[0], m[1] + rhs[1], m[2] + rhs[2], m[3] + rhs[3],
		m[4] + rhs[4], m[5] + rhs[5], m[6] + rhs[6], m[7] + rhs[7],
		m[8] + rhs[8], m[9] + rhs[9], m[10] + rhs[10], m[11] + rhs[11],
		m[12] + rhs[12], m[13] + rhs[13], m[14] + rhs[14], m[15] + rhs[15],
	}
}

func (m Mat4x2) Add(rhs Mat4x2) Mat4x2 {
	return Mat4x2{
		m[0] + rhs[0], m[1] + rhs[1],
		m[2] + rhs[2], m[3] + rhs[3],
		m[4] + rhs[4], m[5] + rhs[5],
		m[6] + rhs[6], m[7] + rhs[7],
	}
}

func (m Mat4x3) Add(rhs Mat4x3) Mat4x3 {
	return Mat4x3{
		m[0] + rhs[0], m[1] + rhs[1], m[2] + rhs[2],
		m[3] + rhs[3], m[4] + rhs[4], m[5] + rhs[5],
		m[6] + rhs[6], m[7] + rhs[7], m[8] + rhs[8],
		m[9] + rhs[9], m[10] + rhs[10], m[11] + rhs[11],
	}
}

func (m Mat4) AddScalar(k float32) Mat4 {
	return Mat4{
		m[0] + k, m[1] + k, m[2] + k, m[3] + k,
		m[4] + k, m[5] + k, m[6] + k, m[7] + k,
		m[8] + k, m[9] + k, m[10] + k, m[11] + k,
		m[12] + k, m[13] + k, m[14] + k, m[15] + k,
	}
}

func (m Mat4x2) AddScalar(k float32) Mat4x2 {
	return Mat4x2{
		m[0] + k, m[1] + k,
		m[2] + k, m[3] + k,
		m[4] + k, m[5] + k,
		m[6] + k, m[7] + k,
	}
}

func (m Mat4x3) AddScalar(k float32) Mat4x3 {
	return Mat4x3{
		m[0] + k, m[1] + k, m[2] + k,
		m[3] + k, m[4] + k, m[5] + k,
		m[6] + k, m[7] + k, m[8] + k,
		m[9] + k, m[10] + k, m[11] + k,
	}
}

func (m Mat4) Sub(rhs Mat4) Mat4 {
	return Mat4{
		m[0] - rhs[0], m[1] - rhs[1], m[2] - rhs[2], m[3] - rhs[3],
		m[4] - rhs[4], m[5] - rhs[5], m[6] - rhs[6], m[7] - rhs[7],
		m[8] - rhs[8], m[9] - rhs[9], m[10] - rhs[10], m[11] - rhs[11],
		m[12] - rhs[12], m[13] - rhs[13], m[14] - rhs[14], m[15] - rhs[15],
	}
}

func (m Mat4x2) Sub(rhs Mat4x2) Mat4x2 {
	return Mat4x2{
		m[0] - rhs[0], m[1] - rhs[1],
		m[2] - rhs[2], m[3] - rhs[3],
		m[4] - rhs[4], m[5] - rhs[5],
		m[6] - rhs[6], m[7] - rhs[7],
	}
}

func (m Mat4x3) Sub(rhs Mat4x3) Mat4x3 {
	return Mat4x3{
		m[0] - rhs[0], m[1] - rhs[1], m[2] - rhs[2],
		m[3] - rhs[3], m[4] - rhs[4], m[5] - rhs[5],
		m[6] - rhs[6], m[7] - rhs[7], m[8] - rhs[8],
		m[9] - rhs[9], m[10] - rhs[10], m[11] - rhs[11],
	}
}

func (m Mat4) SubScalar(k float32) Mat4 {
	return Mat4{
		m[0] - k, m[1] - k, m[2] - k, m[3] - k,
		m[4] - k, m[5] - k, m[6] - k, m[7] - k,
		m[8] - k, m[9] - k, m[10] - k, m[11] - k,
		m[12] - k, m[13] - k, m[14] - k, m[15] - k,
	}
}

func (m Mat4x2) SubScalar(k float32) Mat4x2 {
	return Mat4x2{
		m[0] - k, m[1] - k,
		m[2] - k, m[3] - k,
		m[4] - k, m[5] - k,
		m[6] - k, m[7] - k,
	}
}

func (m Mat4x3) SubScalar(k float32) Mat4x3 {
	return Mat4x3{
		m[0] - k, m[1] - k, m[2] - k,
		m[3] - k, m[4] - k, m[5] - k,
		m[6] - k, m[7] - k, m[8] - k,
		m[9] - k, m[10] - k, m[11] - k,
	}
}

func (m Mat4) MulScalar(k float32) Mat4 {
	return Mat4{
		m[0] * k, m[1] * k, m[2] * k, m[3] * k,
		m[4] * k, m[5] * k, m[6] * k, m[7] * k,
		m[8] * k, m[9] * k, m[10] * k, m[11] * k,
		m[12] * k, m[13] * k, m[14] * k, m[15] * k,
	}
}

func (m Mat4x2) MulScalar(k float32) Mat4x2 {
	return Mat4x2{
		m[0] * k, m[1] * k,
		m[2] * k, m[3] * k,
		m[4] * k, m[5] * k,
		m[6] * k, m[7] * k,
	}
}

func (m Mat4x3) MulScalar(k float32) Mat4x3 {
	return Mat4x3{
		m[0] * k, m[1] * k, m[2] * k,
		m[3] * k, m[4] * k, m[5] * k,
		m[6] * k, m[7] * k, m[8] * k,
		m[9] * k, m[10] * k, m[11] * k,
	}
}

func (m Mat4) Mul4x1(rhs Vec4) Vec4 {
	return Vec4{
		m[0]*rhs[0] + m[1]*rhs[1] + m[2]*rhs[2] + m[3]*rhs[3],
		m[4]*rhs[0] + m[5]*rhs[1] + m[6]*rhs[2] + m[7]*rhs[3],
		m[8]*rhs[0] + m[9]*rhs[1] + m[10]*rhs[2] + m[11]*rhs[3],
		m[12]*rhs[0] + m[13]*rhs[1] + m[14]*rhs[2] + m[15]*rhs[3],
	}
}

func (m Mat4) Mul(rhs Mat4) Mat4 {
	return Mat4{
		m[0]*rhs[0] + m[1]*rhs[4] + m[2]*rhs[8] + m[3]*rhs[12], m[0]*rhs[1] + m[1]*rhs[5] + m[2]*rhs[9] + m[3]*rhs[13], m[0]*rhs[2] + m[1]*rhs[6] + m[2]*rhs[10] + m[3]*rhs[14], m[0]*rhs[3] + m[1]*rhs[7] + m[2]*rhs[11] + m[3]*rhs[15],
		m[4]*rhs[0] + m[5]*rhs[4] + m[6]*rhs[8] + m[7]*rhs[12], m[4]*rhs[1] + m[5]*rhs[5] + m[6]*rhs[9] + m[7]*rhs[13], m[4]*rhs[2] + m[5]*rhs[6] + m[6]*rhs[10] + m[7]*rhs[14], m[4]*rhs[3] + m[5]*rhs[7] + m[6]*rhs[11] + m[7]*rhs[15],
		m[8]*rhs[0] + m[9]*rhs[1] + m[10]*rhs[2] + m[11]*rhs[3], m[8]*rhs[4] + m[9]*rhs[5] + m[10]*rhs[6] + m[11]*rhs[7], m[8]*rhs[8] + m[9]*rhs[9] + m[10]*rhs[10] + m[11]*rhs[11], m[8]*rhs[12] + m[9]*rhs[13] + m[10]*rhs[14] + m[11]*rhs[15],
		m[12]*rhs[0] + m[13]*rhs[1] + m[14]*rhs[2] + m[15]*rhs[3], m[12]*rhs[4] + m[13]*rhs[5] + m[14]*rhs[6] + m[15]*rhs[7], m[12]*rhs[8] + m[13]*rhs[9] + m[14]*rhs[10] + m[15]*rhs[11], m[12]*rhs[12] + m[13]*rhs[13] + m[14]*rhs[14] + m[15]*rhs[15],
	}
}

func (m Mat4) Mul4x2(rhs Mat4x2) Mat4x2 {
	return Mat4x2{
		m[0]*rhs[0] + m[1]*rhs[2] + m[2]*rhs[4] + m[3]*rhs[6], m[0]*rhs[1] + m[1]*rhs[3] + m[2]*rhs[5],
		m[4]*rhs[0] + m[5]*rhs[2] + m[6]*rhs[4] + m[7]*rhs[6], m[4]*rhs[1] + m[5]*rhs[3] + m[6]*rhs[5],
		m[8]*rhs[0] + m[9]*rhs[2] + m[10]*rhs[4] + m[11]*rhs[6], m[8]*rhs[1] + m[9]*rhs[3] + m[10]*rhs[5],
		m[12]*rhs[0] + m[13]*rhs[2] + m[14]*rhs[4] + m[15]*rhs[6], m[12]*rhs[1] + m[13]*rhs[3] + m[14]*rhs[5],
	}
}

func (m Mat4) Mul4x3(rhs Mat4x3) Mat4x3 {
	return Mat4x3{
		m[0]*rhs[0] + m[1]*rhs[3] + m[2]*rhs[6] + m[3]*rhs[9],
		m[4]*rhs[0] + m[5]*rhs[3] + m[6]*rhs[6] + m[7]*rhs[9],
		m[8]*rhs[0] + m[9]*rhs[3] + m[10]*rhs[6] + m[11]*rhs[9],
		m[12]*rhs[0] + m[13]*rhs[3] + m[14]*rhs[6] + m[15]*rhs[9],
	}
}

func Identity4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (m Mat4) ToMat2() Mat2 {
	return Mat2{
		m[0], m[1],
		m[4], m[5],
	}
}

func (m Mat4) ToMat3() Mat3 {
	return Mat3{
		m[0], m[1], m[2],
		m[4], m[5], m[6],
		m[8], m[9], m[10],
	}
}
