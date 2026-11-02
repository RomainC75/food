package shared_domain

import "github.com/google/uuid"

type Uuids interface {
	Generate() uuid.UUID
}
