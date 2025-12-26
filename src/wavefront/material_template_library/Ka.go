package materialtemplatelibrary

import . "scop/src"

type Ka struct {
	Color
}

func NewKa(r, g, b float32) *Ka {
	return &Ka{InitColor(r, g, b, "Ka")}
}
