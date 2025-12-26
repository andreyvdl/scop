package src

import (
	"fmt"
	"strconv"
)

func ParseFloat[T float64 | float32](s string) (T, error) {
	var precision int
	var zero T

	if _, ok := any(zero).(float64); ok {
		precision = 64
	} else {
		precision = 32
	}

	f, e := strconv.ParseFloat(s, precision)
	if e != nil {
		return 0, fmt.Errorf("Error parsing float '%s': %s", s, e)
	}

	return T(f), nil
}
