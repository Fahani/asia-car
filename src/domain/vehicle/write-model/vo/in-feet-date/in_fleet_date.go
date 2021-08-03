package in_fleet_date

import "time"

type InFleetDate struct {
	value time.Time
}

func FromValue(value time.Time) InFleetDate {
	return InFleetDate{value: value}
}

func (n InFleetDate) GetValue() time.Time {
	return n.value
}
