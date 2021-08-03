package fuel_level_test

import (
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/fuel-level"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateValidFuelLevels(t *testing.T) {
	var fuelLevels = []struct {
		expected int
		wanted   error
	}{
		{0, nil},
		{1, nil},
		{99, nil},
		{555555555555555555, nil},
	}

	for _, fl := range fuelLevels {
		fuelLevel, err := fuel_level.FromValue(fl.expected)

		assert.Equal(t, err, fl.wanted)
		assert.Equal(t, fl.expected, fuelLevel.GetValue())
		assert.Nil(t, err)
	}
}

func TestShouldReturnErrorOnInvalidFuelLevels(t *testing.T) {
	var fuelLevels = []struct {
		expected int
		wanted   error
	}{
		{-5, custom_error.EntityNotFound("Invalid Fuel Level -5, the number should be positive integer")},
		{-966, custom_error.EntityNotFound("Invalid Fuel Level -966, the number should be positive integer")},

	}

	for _, fl := range fuelLevels {
		_, err := fuel_level.FromValue(fl.expected)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), fl.wanted.Error())
	}
}
