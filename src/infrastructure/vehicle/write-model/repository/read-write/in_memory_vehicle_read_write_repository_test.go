package in_memory_vehicle_read_write_repository_test

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/write-model/repository/read-write"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestShouldPutAndGetVehicleByChassisNbr(t *testing.T) {
	chassisNbr, _ := chassis_nbr.FromValue("01234567890123456")
	inFleetDate := in_fleet_date.FromValue(time.Now())

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)
	writeRepo := in_memory_vehicle_read_write_repository.NewInMemoryVehicleRepository()

	writeRepo.PutVehicle(*vehicle)

	gotVehicle, err := writeRepo.GetVehicleByChassisNbr(chassisNbr)

	assert.Nil(t, err)

	assert.Equal(t, vehicle.GetInFleetDate().GetValue(), gotVehicle.GetInFleetDate().GetValue())

	assert.Equal(t, vehicle.GetChassisNbr().GetValue(), gotVehicle.GetChassisNbr().GetValue())
}

func TestShouldNotGetANonExistingVehicleByChassisNbr(t *testing.T) {
	chassisNbr, _ := chassis_nbr.FromValue("01234567890123456")

	writeRepo := in_memory_vehicle_read_write_repository.InMemoryVehicleRepository{}

	_, err := writeRepo.GetVehicleByChassisNbr(chassisNbr)

	assert.NotNil(t, err)
	assert.Equal(t, "Vehicle with Chassis Number 01234567890123456 not found", err.Error())
}

func TestShouldPutAndGetVehicleByDeviceSerialNbr(t *testing.T) {
	chassisNbr, _ := chassis_nbr.FromValue("01234567890123456")
	inFleetDate := in_fleet_date.FromValue(time.Now())
	deviceSerialNbr := device_serial_number.FromValue("abcde")

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)
	_ = vehicle.InstallVehicle(deviceSerialNbr)

	writeRepo := in_memory_vehicle_read_write_repository.NewInMemoryVehicleRepository()

	writeRepo.PutVehicle(*vehicle)

	gotVehicle, err := writeRepo.GetVehicleByDeviceSerialNbr(deviceSerialNbr)

	assert.Nil(t, err)

	assert.Equal(t, vehicle.GetInFleetDate().GetValue(), gotVehicle.GetInFleetDate().GetValue())

	assert.Equal(t, vehicle.GetChassisNbr().GetValue(), gotVehicle.GetChassisNbr().GetValue())

	assert.Equal(t, vehicle.GetDeviceSerialNbr().GetValue(), gotVehicle.GetDeviceSerialNbr().GetValue())
}

func TestShouldNotGetANonExistingVehicleByDeviceSerialNbr(t *testing.T) {
	deviceSerialNbr := device_serial_number.FromValue("01234567890123456")

	writeRepo := in_memory_vehicle_read_write_repository.NewInMemoryVehicleRepository()

	_, err := writeRepo.GetVehicleByDeviceSerialNbr(deviceSerialNbr)

	assert.NotNil(t, err)
	assert.Equal(t, "Vehicle with Device Serial Number 01234567890123456 not found", err.Error())
}
