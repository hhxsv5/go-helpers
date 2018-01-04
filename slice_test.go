package go_helpers

import (
	"testing"
)

func TestMakeSliceUnique(t *testing.T) {
	iright := []interface{}{0, 1, 2, 3}
	ints := []interface{}{0, 1, 2, 3, 3, 2, 1}
	nints := MakeSliceUnique(ints)
	if len(nints) != len(iright) {
		t.Error(nints)
	}

	sright := []interface{}{"a", "b", "c"}
	strs := []interface{}{"a", "b", "c", "a", "b", "c", "c"}
	nstrs := MakeSliceUnique(strs)
	if len(nstrs) != len(sright) {
		t.Error(nstrs)
	}
}
