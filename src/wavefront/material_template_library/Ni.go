package materialtemplatelibrary

import "fmt"

type Ni float32

func NewNi(n float32) *Ni {
	if n < 0.001 || n > 10.0 {
		panic(fmt.Sprint("Ni value out of range [0.001, 10.0]: ", n))
	}

	ni := Ni(n)
	return &ni
}
