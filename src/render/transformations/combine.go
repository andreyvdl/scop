package transformations

import (
	. "scop/src/render/matrices"
)

func Combine(t ...Mat4) Mat4 {
	id := Identity4()

	for i := len(t) - 1; i >= 0; i-- { // right2left
		id = id.Mul(t[i])
	}

	return id
}
