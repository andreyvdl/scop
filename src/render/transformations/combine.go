package transformations

import (
	. "scop/src/render/matrices"
)

func Combine(trans Mat4, matrices ...Mat4) Mat4 {
	/*
	 * Note that we first do a translation and then a scale transformation when multiplying matrices.
	 * Matrix multiplication is not commutative, which means their order is important.
	 * When multiplying matrices the right-most matrix is first multiplied with the vector so you should read the multiplications from right to left.
	 */
	for i := len(matrices) - 1; i >= 0; i-- {
		trans = trans.Mul(matrices[i])
	}

	return trans
}
