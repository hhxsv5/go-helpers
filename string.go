package go_helpers

import (
	"strconv"
)

func StrToInt64(s string) int64 {
	r, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		return 0
	}
	return r
}

func StrToFloat64(s string) float64 {
	r, e := strconv.ParseFloat(s, 64)
	if e != nil {
		return 0
	}
	return r
}
