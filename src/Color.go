package src

import (
	"fmt"
)

type Color struct {
	r, g, b float32
}

func InitColor(r, g, b float32, origin string) Color {
	if r < 0.0 || r > 1.0 || g < 0.0 || g > 1.0 || b < 0.0 || b > 1.0 {
		panic(fmt.Sprintf("%s: Color value out of range [0.0, 1.0]: %f %f %f", origin, r, g, b))
	}

	return Color{r, g, b}
}

func Str2RGB(values []string) (r, g, b float32) {
	r, err := ParseFloat[float32](values[0])
	if err != nil {
		panic(err)
	}

	g, err = ParseFloat[float32](values[1])
	if err != nil {
		panic(err)
	}

	b, err = ParseFloat[float32](values[2])
	if err != nil {
		panic(err)
	}

	return
}
