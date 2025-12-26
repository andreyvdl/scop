package materialtemplatelibrary

import "fmt"

type Ns float32

func NewNs(n float32) *Ns {
	if n < 0.0 || n > 1000.0 {
		panic(fmt.Sprint("Ns value out of range [0.0, 1000.0]: ", n))
	}

	ns := Ns(n)
	return &ns
}
