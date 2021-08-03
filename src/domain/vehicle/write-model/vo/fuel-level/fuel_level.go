package fuel_level

import (
	"fmt"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
)

type FuelLevel struct {
	value int
}

func FromValue(value int) (FuelLevel, error) {
	err := assertValidValue(value)
	if err != nil {
		return FuelLevel{}, err
	}
	return FuelLevel{value: value}, err
}

func assertValidValue(value int) error {

	if value < 0 {
		return custom_error.EntityNotFound(fmt.Sprintf("Invalid Fuel Level %d, the number should be positive integer", value))
	}
	return nil
}

func (n FuelLevel) GetValue() int {
	return n.value
}
