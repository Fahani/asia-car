package battery_level_test

import (
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/battery-level"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateValidBatteryLevels(t *testing.T) {
	var batteryLevels = []struct {
		expected int
		wanted   error
	}{
		{0, nil},
		{1, nil},
		{99, nil},
		{555555555555555555, nil},
	}

	for _, bl := range batteryLevels {
		batteryLevel, err := battery_level.FromValue(bl.expected)
		assert.Nil(t, err)
		assert.Equal(t, bl.wanted, err)
		assert.Equal(t, bl.expected, batteryLevel.GetValue())
	}
}

func TestShouldReturnErrorOnInvalidBatteryLevels(t *testing.T) {
	var batteryLevels = []struct {
		expected int
		wanted   error
	}{
		{-5, custom_error.InvalidArgument("Invalid Battery Level -5, the number should be positive integer")},
		{-966, custom_error.InvalidArgument("Invalid Battery Level -966, the number should be positive integer")},
	}

	for _, lp := range batteryLevels {
		_, err := battery_level.FromValue(lp.expected)

		assert.NotNil(t, err)
		assert.Equal(t, lp.wanted.Error(), err.Error())
	}
}
