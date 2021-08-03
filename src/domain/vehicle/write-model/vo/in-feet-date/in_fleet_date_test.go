package in_fleet_date_test

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestShouldCreateValidInFleetDates(t *testing.T) {
	var inFleetDates = []struct {
		expected time.Time
		wanted   error
	}{
		{time.Now(), nil},
		{time.Date(1990, time.May, 10, 23, 12, 5, 3, time.UTC), nil},
	}

	for _, ifd := range inFleetDates {
		inFleetDate := in_fleet_date.FromValue(ifd.expected)

		assert.Equal(t, inFleetDate.GetValue(), ifd.expected)
	}
}
