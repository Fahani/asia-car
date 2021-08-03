package chassis_nbr_test

import (
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateValidChassisNumbers(t *testing.T) {
	var chassisNumbers = []struct {
		expected string
		wanted   error
	}{
		{"00000000000000000", nil},
		{"01234567890123456", nil},
		{"1a1a1aa11a1a1a1a1", nil},
		{"abcdefghijklmnopq", nil},
	}

	for _, cn := range chassisNumbers {
		chassisNbr, err := chassis_nbr.FromValue(cn.expected)

		assert.Nil(t, err)
		assert.Equal(t, cn.wanted, err)
		assert.Equal(t, cn.expected, chassisNbr.GetValue())
	}
}

func TestShouldReturnErrorOnInvalidChassisNumbers(t *testing.T) {
	var chassisNumbers = []struct {
		expected string
		wanted   error
	}{
		{"0123", custom_error.InvalidArgument("Invalid Chassis Number 0123, the number should be 17 digits alphanumeric")},
		{"012345678901234567", custom_error.InvalidArgument("Invalid Chassis Number 012345678901234567, the number should be 17 digits alphanumeric")},
		{"abcdefghijklmnopñ", custom_error.InvalidArgument("Invalid Chassis Number abcdefghijklmnopñ, the number should be 17 digits alphanumeric")},
	}

	for _, cn := range chassisNumbers {
		_, err := chassis_nbr.FromValue(cn.expected)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), cn.wanted.Error())
	}
}
