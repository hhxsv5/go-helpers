package go_helpers

import (
	"testing"
)

func TestMakeBoolSliceUnqiue(t *testing.T) {
	r := []bool{true, false}
	a := []bool{true, false, false, true}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
	for _, v := range b {
		t.Log(v.(bool))
	}
}

func TestMakeInt32SliceUnique(t *testing.T) {
	r := []int32{0, 1, 2, 3}
	a := []int32{0, 1, 2, 3, 3, 2, 1}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
	for _, v := range b {
		t.Log(v.(int32))
	}
}

func TestMakeInt64SliceUnique(t *testing.T) {
	r := []int64{0, 1, 2, 3}
	a := []int64{0, 1, 2, 3, 3, 2, 1}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
	for _, v := range b {
		t.Log(v.(int64))
	}
}

func TestMakeFloat32SliceUnique(t *testing.T) {
	r := []float32{0.01, 1.22, 0.0}
	a := []float32{0, 0.0, 1.22, 0.01, 0.01, 1.22, 0.0, 0}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
	for _, v := range b {
		t.Log(v.(float32))
	}
}

func TestMakeFloat64SliceUnique(t *testing.T) {
	r := []float64{0.001, 1.2222, 0.000, 0.01}
	a := []float64{0, 0.00, 1.2222, 0.01, 0.001, 1.2222, 0.0, 0}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
	for _, v := range b {
		t.Log(v.(float64))
	}
}

func TestMakeStringSliceUnqiue(t *testing.T) {
	r := []string{"a", "b", "c"}
	a := []string{"a", "b", "c", "a", "b", "c", "c"}
	b := MakeSliceUnique(a)
	t.Log(b)
	if len(r) != len(b) {
		t.Error(b, r)
	}
	for _, v := range b {
		t.Log(v.(string))
	}
}

func TestStrPos(t *testing.T) {
	strs := []string{"a1", "b2", "c3", "d4"}
	target := "c3"
	pos := StrPos(strs, target)
	if pos != 2 {
		t.Error(pos)
	}
	t.Log(pos)
}

func TestInt64Pos(t *testing.T) {
	ints := []int64{-100, 100, 200, 0, 300}
	target := int64(0)
	pos := Int64Pos(ints, target)
	if pos != 3 {
		t.Error(pos)
	}
	t.Log(pos)
}
