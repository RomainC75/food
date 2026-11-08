package shared_infra

import "time"

type InMemDateTime struct {
	now time.Time
}

func NewInMemDateTime() *InMemDateTime {
	return &InMemDateTime{
		now: time.Now(),
	}
}

func (imdt *InMemDateTime) Now() time.Time {
	return imdt.now
}

func (imdt *InMemDateTime) SetNow(now time.Time) *InMemDateTime {
	imdt.now = now
	return imdt
}
