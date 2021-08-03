package battery_level

import (
	"fmt"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
)

type BatteryLevel struct {
	value int
}

func FromValue(value int) (BatteryLevel, error) {
	err := assertValidValue(value)
	if err != nil {
		return BatteryLevel{}, err
	}
	return BatteryLevel{value: value}, err
}

func assertValidValue(value int) error {

	if value < 0 {
		return custom_error.InvalidArgument(fmt.Sprintf("Invalid Battery Level %d, the number should be positive integer", value))
	}
	return nil
}

func (n BatteryLevel) GetValue() int {
	return n.value
}
