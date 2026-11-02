package shared_infra

import "github.com/google/uuid"

type UuidGenerator struct {
}

func NewUuidGenerator() *UuidGenerator {
	return &UuidGenerator{}
}

func (ug *UuidGenerator) Generate() uuid.UUID {
	return uuid.New()
}
