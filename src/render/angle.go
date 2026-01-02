package render

import "math"

func Deg2Rad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

func Rad2Deg(rad float64) float64 {
	return rad * 180.0 / math.Pi
}
