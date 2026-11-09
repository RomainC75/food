package shared_infra

import "github.com/google/uuid"

type InMemUuidGen struct {
	expected     []uuid.UUID
	currentIndex int
}

func NewInMemUuidGen() *InMemUuidGen {
	return &InMemUuidGen{}
}

func (imu *InMemUuidGen) SetNext(id uuid.UUID) *InMemUuidGen {
	imu.expected = []uuid.UUID{id}
	imu.currentIndex = 0
	return imu
}

func (imu *InMemUuidGen) SetMultipleUuids(uuids []uuid.UUID) *InMemUuidGen {
	imu.expected = uuids
	imu.currentIndex = 0
	return imu
}

func (imu *InMemUuidGen) Generate() uuid.UUID {
	if imu.currentIndex == len(imu.expected) {
		panic("InMemUuidGenerator : no more uuids")
	}

	id := imu.expected[imu.currentIndex]
	imu.currentIndex++
	return id
}
