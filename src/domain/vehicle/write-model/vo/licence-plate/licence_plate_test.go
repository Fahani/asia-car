package licence_plate_test

import (
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/licence-plate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateValidLicencePlates(t *testing.T) {
	var licencePlates = []struct {
		expected string
		wanted   error
	}{
		{"0", nil},
		{"123", nil},
		{"1a1a", nil},
		{"abc", nil},
	}

	for _, lp := range licencePlates {
		licencePlate, err := licence_plate.FromValue(lp.expected)

		assert.Nil(t, err)
		assert.Equal(t, err, lp.wanted)
		assert.Equal(t, licencePlate.GetValue(), lp.expected)
	}
}

func TestShouldReturnErrorOnInvalidLicencePlates(t *testing.T) {
	var licencePlates = []struct {
		expected string
		wanted   error
	}{
		{"", custom_error.InvalidArgument("Invalid Licence Plate , the number should be at least one digit alphanumeric")},
		{"_", custom_error.InvalidArgument("Invalid Licence Plate _, the number should be at least one digit alphanumeric")},

	}

	for _, lp := range licencePlates {
		_, err := licence_plate.FromValue(lp.expected)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), lp.wanted.Error())
	}
}
