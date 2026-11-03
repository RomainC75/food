package recipe_dom

import "errors"

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
	return Quantity{
		value: v,
		unit:  u,
	}
}
