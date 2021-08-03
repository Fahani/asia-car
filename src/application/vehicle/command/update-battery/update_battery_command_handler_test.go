package update_battery_command_test

import (
	"github.com/fahani/asia-car/src/application/vehicle/command/update-battery"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/battery-level"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createInFleetDateVOChassisNumberVODeviceSerialNumberVOBatteryVO(inFleetDate, chassisNumber, deviceSerialNumber string, batteryLevel int) (in_fleet_date.InFleetDate, chassis_nbr.ChassisNbr, device_serial_number.DeviceSerialNumber, battery_level.BatteryLevel, error) {
	chassisNumberVO, err := chassis_nbr.FromValue(chassisNumber)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, err
	}

	t, err := time.Parse(time.RFC3339, inFleetDate)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, err
	}

	batteryLevelVO, err := battery_level.FromValue(batteryLevel)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, err
	}

	inFleetDateVO := in_fleet_date.FromValue(t)

	deviceSerialNumberVO := device_serial_number.FromValue(deviceSerialNumber)

	return inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, batteryLevelVO, nil
}

func TestShouldUpdateBatteryLevelCorrectly(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	batteryLevel := 50

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, batterLevelVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOBatteryVO(inFleetDate, chassisNumber, deviceSerialNumber, batteryLevel)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)
	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	err = vehicle.SetBatteryLevel(batterLevelVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)
	vehicleWriteRepositoryMocked.EXPECT().PutVehicle(*vehicle).Times(1)

	updateBatteryLevelCommand := update_battery_command.NewUpdateBatteryCommand(deviceSerialNumber, batteryLevel)

	updateBatteryLevelCommandHandler := update_battery_command.NewUpdateBatteryLevelCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateBatteryLevelCommandHandler.Handle(updateBatteryLevelCommand)

	assert.Nil(t, err)
}

func TestShouldReturnErrorIfVehicleIsNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	batteryLevel := 50

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, batterLevelVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOBatteryVO(inFleetDate, chassisNumber, deviceSerialNumber, batteryLevel)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)
	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, custom_error.EntityNotFound("Vehicle with Device Serial Number abc not found")).Times(1)

	err = vehicle.SetBatteryLevel(batterLevelVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)


	updateBatteryLevelCommand := update_battery_command.NewUpdateBatteryCommand(deviceSerialNumber, batteryLevel)

	updateBatteryLevelCommandHandler := update_battery_command.NewUpdateBatteryLevelCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateBatteryLevelCommandHandler.Handle(updateBatteryLevelCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Vehicle with Device Serial Number abc not found", err.Error())
}

func TestShouldReturnErrorIfBatteryLevelIsNegative(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	batteryLevel := 50

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, batterLevelVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOBatteryVO(inFleetDate, chassisNumber, deviceSerialNumber, batteryLevel)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)
	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	err = vehicle.SetBatteryLevel(batterLevelVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)


	updateBatteryLevelCommand := update_battery_command.NewUpdateBatteryCommand(deviceSerialNumber, -10)

	updateBatteryLevelCommandHandler := update_battery_command.NewUpdateBatteryLevelCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateBatteryLevelCommandHandler.Handle(updateBatteryLevelCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Battery Level -10, the number should be positive integer", err.Error())
}

func TestShouldReturnErrorIfVehicleIsNotInstalled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	batteryLevel := 50

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, _, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOBatteryVO(inFleetDate, chassisNumber, deviceSerialNumber, batteryLevel)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	updateBatteryLevelCommand := update_battery_command.NewUpdateBatteryCommand(deviceSerialNumber, batteryLevel)

	updateBatteryLevelCommandHandler := update_battery_command.NewUpdateBatteryLevelCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateBatteryLevelCommandHandler.Handle(updateBatteryLevelCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't set the battery level of the vehicle if it hasn't been in fleet before", err.Error())
}