package device_serial_number_test

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateValidDeviceSerialNumbers(t *testing.T) {
	var deviceSerialNumbers = []struct {
		expected string
		wanted   error
	}{
		{"", nil},
		{"123", nil},
		{"1a1a", nil},
		{"abc", nil},
		{"abc-e", nil},
	}

	for _, dsn := range deviceSerialNumbers {
		device := device_serial_number.FromValue(dsn.expected)

		assert.Equal(t, dsn.expected, device.GetValue())
	}
}
