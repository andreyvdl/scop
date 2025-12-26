package wavefront

type Normal struct {
	X, Y, Z float32
}

func NewNormal(x, y, z float32) *Normal {
	return &Normal{x, y, z}
}
