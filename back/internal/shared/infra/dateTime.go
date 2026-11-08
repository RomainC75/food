package shared_infra

import "time"

type DateTime struct {
}

func NewDateTime() *DateTime {
	return &DateTime{}
}

func (imdt *DateTime) Now() time.Time {
	return time.Now()
}
