package get_telemetries_by_chassis_number_query

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"time"
)

type GetTelemetriesByChassisNumberQueryResponse struct {
	ChassisNumber      string    `json:"chassis_nbr"`
	LicensePlate       string    `json:"license_plate"`
	Brand              string    `json:"brand"`
	Category           string    `json:"category"`
	InFleetDate        time.Time `json:"in_fleet_date"`
	DeviceSerialNumber string    `json:"device_serial_number"`
	BatteryLevel       int       `json:"battery_level"`
	FuelLevel          int       `json:"fuel_level"`
	CurrentMileage     int       `json:"current_mileage"`
}

func NewGetTelemetriesBySerialNumberQueryResponse(vehicle vehicle_aggregate.Vehicle) GetTelemetriesByChassisNumberQueryResponse {
	return GetTelemetriesByChassisNumberQueryResponse{
		ChassisNumber:      vehicle.GetChassisNbr().GetValue(),
		LicensePlate:       vehicle.GetLicencePlate().GetValue(),
		Brand:              vehicle.GetBrand().GetValue(),
		Category:           vehicle.GetCategory().GetValue(),
		InFleetDate:        vehicle.GetInFleetDate().GetValue(),
		DeviceSerialNumber: vehicle.GetDeviceSerialNbr().GetValue(),
		BatteryLevel:       vehicle.GetBatteryLevel().GetValue(),
		FuelLevel:          vehicle.GetFuel().GetValue(),
		CurrentMileage:     vehicle.GetMileage().GetValue(),
	}
}