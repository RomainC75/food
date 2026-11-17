package recipe_dom

import (
	"encoding/json"
	"errors"
)

var (
	ErrNegativeValue = "negative quatity value"
)

type Quantity struct {
	value float32
	unit  BaseUnit
}

func NewQuantity(v float32, u BaseUnit) (Quantity, error) {
	if v < 0 {
		return Quantity{}, errors.New(ErrNegativeValue)
	}
	return Quantity{
		value: v,
		unit:  u,
	}, nil
}

func MustNewQuantity(v float32, u BaseUnit) Quantity {
	if v < 0 {
		panic("quantity cannot be negative")
	}
	return Quantity{
		value: v,
		unit:  u,
	}
}

func (q *Quantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float32  `json:"value"`
		Unit  BaseUnit `json:"unit"`
	}{
		Value: q.value,
		Unit:  q.unit,
	})
}
