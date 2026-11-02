package shared_infra

import "github.com/google/uuid"

type InMemUuidGen struct {
	expected uuid.UUID
}

func NewInMemUuidGen() *InMemUuidGen {
	return &InMemUuidGen{}
}

func (imu *InMemUuidGen) AddExpected(id uuid.UUID) *InMemUuidGen {
	imu.expected = id
	return imu
}

func (imu *InMemUuidGen) Generate() uuid.UUID {
	return imu.expected
}
