package shared_infra

import "github.com/google/uuid"

type InMemUuid struct {
	expected uuid.UUID
}

func NewInMemUuidGen() *InMemUuid {
	return &InMemUuid{}
}

func (imu *InMemUuid) AddExpected(id uuid.UUID) *InMemUuid {
	imu.expected = id
	return imu
}

func (imu *InMemUuid) Generate() uuid.UUID {
	return imu.expected
}
