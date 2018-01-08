package go_helpers

import "testing"

func TestNowDate(t *testing.T) {
	nd := NowDate()
	t.Log(nd)
}

func TestNowDateTime(t *testing.T) {
	ndt := NowDateTime()
	t.Log(ndt)
}

func TestParseStringTime(t *testing.T) {
	tm, err := ParseStringTime("2018-01-08 14:23:00", "Asia/Shanghai")
	if err != nil {
		t.Error(err)
	}
	t.Log(tm.String())
}
