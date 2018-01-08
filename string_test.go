package go_helpers

import "testing"

func TestStrToInt64(t *testing.T) {
	i := StrToInt64("3000")
	if i != 3000 {
		t.Error(i)
	}
	t.Log(i)

}

func TestStrToFloat64(t *testing.T) {
	i := StrToFloat64("3000.1234")
	if i != 3000.1234 {
		t.Error(i)
	}
	t.Log(i)
}
