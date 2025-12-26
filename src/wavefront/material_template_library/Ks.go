package materialtemplatelibrary

import . "scop/src"

type Ks struct {
	Color
}

func NewKs(r, g, b float32) *Ks {
	return &Ks{InitColor(r, g, b, "Ks")}
}
