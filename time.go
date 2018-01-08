package go_helpers

import "time"

const (
	TimeLayoutDate     = "2006-01-02"
	TimeLayoutDateTime = "2006-01-02 15:04:05"
)

func NowUnix() int64 {
	return time.Now().Unix()
}

func NowDate() string {
	return time.Now().Format(TimeLayoutDate)
}

func NowDateTime() string {
	return time.Now().Format(TimeLayoutDateTime)
}

func ParseDate(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDate, dt)
}

func ParseDateTime(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDateTime, dt)
}

func ParseStringTime(tm, lc string) (time.Time, error) {
	loc, err := time.LoadLocation(lc)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(TimeLayoutDateTime, tm, loc)
}
