package shared_domain

import "time"

type DateTime interface {
	Now() time.Time
}
