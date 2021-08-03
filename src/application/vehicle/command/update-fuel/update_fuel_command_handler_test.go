package update_fuel_command_test

import (
	"github.com/fahani/asia-car/src/application/vehicle/command/update-fuel"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/fuel-level"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createInFleetDateVOChassisNumberVODeviceSerialNumberVOFuelVO(inFleetDate, chassisNumber, deviceSerialNumber string, fuel int) (in_fleet_date.InFleetDate, chassis_nbr.ChassisNbr, device_serial_number.DeviceSerialNumber, fuel_level.FuelLevel, error) {
	chassisNumberVO, err := chassis_nbr.FromValue(chassisNumber)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, fuel_level.FuelLevel{}, err
	}

	t, err := time.Parse(time.RFC3339, inFleetDate)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, fuel_level.FuelLevel{}, err
	}

	inFleetDateVO := in_fleet_date.FromValue(t)

	deviceSerialNumberVO := device_serial_number.FromValue(deviceSerialNumber)

	fuelVO, err := fuel_level.FromValue(fuel)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, fuel_level.FuelLevel{}, err
	}

	return inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, fuelVO, nil
}

func TestShouldIncrementFuelCorrectly(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	fuel := 50
	updateType := "increment"

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, fuelVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOFuelVO(inFleetDate, chassisNumber, deviceSerialNumber, fuel)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)
	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	err = vehicle.IncreaseFuel(fuelVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)
	vehicleWriteRepositoryMocked.EXPECT().PutVehicle(*vehicle).Times(1)

	updateFuelCommand := update_fuel_command.NewUpdateFuelCommand(fuel, deviceSerialNumber, updateType)

	updateFuelCommandHandler := update_fuel_command.NewUpdateFuelCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateFuelCommandHandler.Handle(updateFuelCommand)

	assert.Nil(t, err)
}

func TestShouldDecrementFuelCorrectly(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	fuel := 50
	updateType := "decrement"

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, fuelVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOFuelVO(inFleetDate, chassisNumber, deviceSerialNumber, fuel)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	err = vehicle.InstallVehicle(deviceSerialNumberVO)
	assert.Nil(t, err)

	err = vehicle.IncreaseFuel(fuelVO)
	assert.Nil(t, err)

	err = vehicle.IncreaseFuel(fuelVO)
	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	err = vehicle.DecreaseFuel(fuelVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)
	vehicleWriteRepositoryMocked.EXPECT().PutVehicle(*vehicle).Times(1)

	updateFuelCommand := update_fuel_command.NewUpdateFuelCommand(fuel, deviceSerialNumber, updateType)

	updateFuelCommandHandler := update_fuel_command.NewUpdateFuelCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateFuelCommandHandler.Handle(updateFuelCommand)

	assert.Nil(t, err)
}

func TestShouldReturnAnErrorIfVehicleNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	fuel := 50
	updateType := "decrement"

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, fuelVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOFuelVO(inFleetDate, chassisNumber, deviceSerialNumber, fuel)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	err = vehicle.InstallVehicle(deviceSerialNumberVO)
	assert.Nil(t, err)

	err = vehicle.IncreaseFuel(fuelVO)
	assert.Nil(t, err)

	err = vehicle.IncreaseFuel(fuelVO)
	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, custom_error.EntityNotFound("Vehicle with Device Serial Number abc not found")).Times(1)

	err = vehicle.DecreaseFuel(fuelVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	updateFuelCommand := update_fuel_command.NewUpdateFuelCommand(fuel, deviceSerialNumber, updateType)

	updateFuelCommandHandler := update_fuel_command.NewUpdateFuelCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateFuelCommandHandler.Handle(updateFuelCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Vehicle with Device Serial Number abc not found", err.Error())
}

func TestShouldReturnAnErrorIfFuelIsNegative(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	fuel := 50
	updateType := "decrement"

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, fuelVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOFuelVO(inFleetDate, chassisNumber, deviceSerialNumber, fuel)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	err = vehicle.InstallVehicle(deviceSerialNumberVO)
	assert.Nil(t, err)

	err = vehicle.IncreaseFuel(fuelVO)
	assert.Nil(t, err)

	err = vehicle.IncreaseFuel(fuelVO)
	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	err = vehicle.DecreaseFuel(fuelVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	updateFuelCommand := update_fuel_command.NewUpdateFuelCommand(-10, deviceSerialNumber, updateType)

	updateFuelCommandHandler := update_fuel_command.NewUpdateFuelCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateFuelCommandHandler.Handle(updateFuelCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Fuel Level -10, the number should be positive integer", err.Error())
}

func TestShouldReturnAnErrorIfDecreasingFuelResultInANegativeOperation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	fuel := 50
	updateType := "decrement"

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, _, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOFuelVO(inFleetDate, chassisNumber, deviceSerialNumber, fuel)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	err = vehicle.InstallVehicle(deviceSerialNumberVO)
	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	updateFuelCommand := update_fuel_command.NewUpdateFuelCommand(fuel, deviceSerialNumber, updateType)

	updateFuelCommandHandler := update_fuel_command.NewUpdateFuelCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateFuelCommandHandler.Handle(updateFuelCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Fuel Level -50, the number should be positive integer", err.Error())
}