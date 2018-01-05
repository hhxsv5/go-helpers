package go_helpers

import (
	"testing"
)

func TestIntMakeSliceUnique(t *testing.T) {
	iright := []interface{}{0, 1, 2, 3}
	ints := []interface{}{0, 1, 2, 3, 3, 2, 1}
	nints := MakeSliceUnique(ints)
	if len(nints) != len(iright) {
		t.Error(nints)
	}
}

func TestStringMakeSliceUnique(t *testing.T) {
	iright := []interface{}{0, 1, 2, 3}
	ints := []interface{}{0, 1, 2, 3, 3, 2, 1}
	nints := MakeSliceUnique(ints)
	if len(nints) != len(iright) {
		t.Error(nints)
	}
}
