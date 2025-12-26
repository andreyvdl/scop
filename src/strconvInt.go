package src

import (
	"fmt"
	"strconv"
)

func ParseUint[T uint | uint8 | uint16 | uint32 | uint64](s string) (T, error) {
	var precision int
	var zero T

	switch any(zero).(type) {
	case uint64:
		precision = 64

	case uint32, uint:
		precision = 32

	case uint16:
		precision = 16

	case uint8:
		precision = 8
	}

	f, e := strconv.ParseUint(s, 10, precision)
	if e != nil {
		return 0, fmt.Errorf("Error parsing uint '%s': %s", s, e)
	}

	return T(f), nil
}

func ParseInt[T int | int8 | int16 | int32 | int64](s string) (T, error) {
	var precision int
	var zero T

	switch any(zero).(type) {
	case int64:
		precision = 64

	case int32, int:
		precision = 32

	case int16:
		precision = 16

	case int8:
		precision = 8
	}

	f, e := strconv.ParseInt(s, 10, precision)
	if e != nil {
		return 0, fmt.Errorf("Error parsing int '%s': %s", s, e)
	}

	return T(f), nil
}
