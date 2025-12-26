package materialtemplatelibrary

import . "scop/src"

type Kd struct {
	Color
}

func NewKd(r, g, b float32) *Kd {
	return &Kd{InitColor(r, g, b, "Kd")}
}
