package install_vehicle_command_test

import (
	"github.com/fahani/asia-car/src/application/vehicle/command/install-vehicle"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createInFleetDateVOChassisNumberVODeviceSerialNumberVO(inFleetDate, chassisNumber, deviceSerialNumber string) (in_fleet_date.InFleetDate, chassis_nbr.ChassisNbr, device_serial_number.DeviceSerialNumber, error) {
	chassisNumberVO, err := chassis_nbr.FromValue(chassisNumber)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, err
	}

	t, err := time.Parse(time.RFC3339, inFleetDate)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, err
	}

	inFleetDateVO := in_fleet_date.FromValue(t)

	deviceSerialNumberVO := device_serial_number.FromValue(deviceSerialNumber)

	return inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, nil
}

func TestShouldInstallVehicleCorrectly(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVO(inFleetDate, chassisNumber, deviceSerialNumber)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByChassisNbr(chassisNumberVO).Return(*vehicle, nil).Times(1)

	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)
	vehicleWriteRepositoryMocked.EXPECT().PutVehicle(*vehicle).Times(1)

	installCommand := install_vehicle_command.NewInstallVehicleCommand(deviceSerialNumber, chassisNumber)

	installCommandHandler := install_vehicle_command.NewInstallVehicleCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = installCommandHandler.Handle(installCommand)

	assert.Nil(t, err)
}

func TestShouldReturnErrorIfChassisNumberIsNotValid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVO(inFleetDate, chassisNumber, deviceSerialNumber)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)

	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	chassisNumber = "0123456789012345"

	installCommand := install_vehicle_command.NewInstallVehicleCommand(deviceSerialNumber, chassisNumber)

	installCommandHandler := install_vehicle_command.NewInstallVehicleCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = installCommandHandler.Handle(installCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Chassis Number 0123456789012345, the number should be 17 digits alphanumeric", err.Error())
}

func TestShouldReturnErrorWhenNotFindingVehicle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVO(inFleetDate, chassisNumber, deviceSerialNumber)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByChassisNbr(chassisNumberVO).Return(*vehicle, custom_error.EntityNotFound("Vehicle with Chassis Number 01234567890123456 not found")).Times(1)

	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	installCommand := install_vehicle_command.NewInstallVehicleCommand(deviceSerialNumber, chassisNumber)

	installCommandHandler := install_vehicle_command.NewInstallVehicleCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = installCommandHandler.Handle(installCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Vehicle with Chassis Number 01234567890123456 not found", err.Error())
}

func TestShouldReturnErrorWhenVehicleIsNotInFleet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"

	_, chassisNumberVO, _, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVO(inFleetDate, chassisNumber, deviceSerialNumber)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.Vehicle{}

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByChassisNbr(chassisNumberVO).Return(vehicle, nil).Times(1)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	installCommand := install_vehicle_command.NewInstallVehicleCommand(deviceSerialNumber, chassisNumber)

	installCommandHandler := install_vehicle_command.NewInstallVehicleCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = installCommandHandler.Handle(installCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't install a vehicle if hasn't been in fleet before", err.Error())
}

