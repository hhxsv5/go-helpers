package go_helpers

import (
	"testing"
)

func TestMakeBoolSliceUnqiue(t *testing.T) {
	r := []bool{true, false}
	a := []bool{true, false, false, true}
	b := MakeSliceUnique(a)
	if len(r) != len(b) {
		t.Error(b)
	}
}

func TestMakeIntSliceUnique(t *testing.T) {
	r := []int{0, 1, 2, 3}
	a := []int{0, 1, 2, 3, 3, 2, 1}
	b := MakeSliceUnique(a)
	if len(r) != len(b) {
		t.Error(b)
	}
}

func TestMakeFloatSliceUnique(t *testing.T) {
	r := []float32{0.01, 1.22, 0.0}
	a := []float32{0, 0.0, 1.22, 0.01, 0.01, 1.22, 0.0, 0}
	b := MakeSliceUnique(a)
	if len(r) != len(b) {
		t.Error(b)
	}
}

func TestMakeStringSliceUnqiue(t *testing.T) {
	r := []string{"a", "b", "c"}
	a := []string{"a", "b", "c", "a", "b", "c", "c"}
	b := MakeSliceUnique(a)
	if len(r) != len(b) {
		t.Error(b)
	}
}
