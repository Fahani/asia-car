package current_mileage_test

import (
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/current-mileage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateValidCurrentMileages(t *testing.T) {
	var currentMileages = []struct {
		expected int
		wanted   error
	}{
		{0, nil},
		{1, nil},
		{99, nil},
		{555555555555555555, nil},
	}

	for _, cm := range currentMileages {
		currentMileage, err := current_mileage.FromValue(cm.expected)

		assert.Nil(t, err)
		assert.Equal(t, cm.wanted, err)
		assert.Equal(t, cm.expected, currentMileage.GetValue())
	}
}

func TestShouldReturnErrorOnInvalidCurrentMileages(t *testing.T) {
	var currentMileages = []struct {
		expected int
		wanted   error
	}{
		{-5, custom_error.InvalidArgument("Invalid Current Mileage -5, the number should be positive integer")},
		{-966, custom_error.InvalidArgument("Invalid Current Mileage -966, the number should be positive integer")},
	}

	for _, cm := range currentMileages {
		_, err := current_mileage.FromValue(cm.expected)

		assert.NotNil(t, err)
		assert.Equal(t, cm.wanted.Error(), err.Error())
	}
}
