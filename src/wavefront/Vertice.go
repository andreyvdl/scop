package wavefront

import "scop/src"

// Represents a vertice
// x, y and z are mandatory values
// w is optinal, by default 1.0
type Vertex struct {
	X, Y, Z, W float32
}

type vertexOption func(*Vertex)

func NewVertice(x, y, z float32, opts ...vertexOption) *Vertex {
	v := &Vertex{x, y, z, 1.0}

	for _, opt := range opts {
		opt(v)
	}

	return v
}

func WithW(w float32) vertexOption {
	return func(v *Vertex) {
		v.W = w
	}
}

func Str2Vertex(values []string) (x, y, z float32) {
	var err error

	x, err = src.ParseFloat[float32](values[0])
	if err != nil {
		panic(err)
	}

	y, err = src.ParseFloat[float32](values[1])
	if err != nil {
		panic(err)
	}

	z, err = src.ParseFloat[float32](values[2])
	if err != nil {
		panic(err)
	}

	return
}
