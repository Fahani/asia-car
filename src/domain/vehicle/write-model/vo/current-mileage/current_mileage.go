package current_mileage

import (
	"fmt"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
)

type CurrentMileage struct {
	value int
}

func FromValue(value int) (CurrentMileage, error) {
	err := assertValidValue(value)
	if err != nil {
		return CurrentMileage{}, err
	}
	return CurrentMileage{value: value}, err
}

func assertValidValue(value int) error {

	if value < 0 {
		return custom_error.InvalidArgument(fmt.Sprintf("Invalid Current Mileage %d, the number should be positive integer", value))
	}
	return nil
}

func (n CurrentMileage) GetValue() int {
	return n.value
}
