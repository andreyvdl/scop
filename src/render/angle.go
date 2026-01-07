package render

import "math"

const (
	RAD = 0.01745329
)

func Deg2Rad(deg float64) float64 {
	return deg * RAD
}

func Rad2Deg(rad float64) float64 {
	return rad * 180.0 / math.Pi
}
