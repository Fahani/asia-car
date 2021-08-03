package update_mileage_command_test

import (
	"github.com/fahani/asia-car/src/application/vehicle/command/update-mileage"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/current-mileage"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createInFleetDateVOChassisNumberVODeviceSerialNumberVOMileageVO(inFleetDate, chassisNumber, deviceSerialNumber string, mileage int) (in_fleet_date.InFleetDate, chassis_nbr.ChassisNbr, device_serial_number.DeviceSerialNumber, current_mileage.CurrentMileage, error) {
	chassisNumberVO, err := chassis_nbr.FromValue(chassisNumber)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, current_mileage.CurrentMileage{}, err
	}

	t, err := time.Parse(time.RFC3339, inFleetDate)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, current_mileage.CurrentMileage{}, err
	}

	mileageVo, err := current_mileage.FromValue(mileage)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, device_serial_number.DeviceSerialNumber{}, current_mileage.CurrentMileage{}, err
	}

	inFleetDateVO := in_fleet_date.FromValue(t)

	deviceSerialNumberVO := device_serial_number.FromValue(deviceSerialNumber)

	return inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, mileageVo, nil
}

func TestShouldUpdateMileageCorrectly(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	mileage := 50

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, mileageVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOMileageVO(inFleetDate, chassisNumber, deviceSerialNumber, mileage)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)
	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	err = vehicle.SetMileage(mileageVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)
	vehicleWriteRepositoryMocked.EXPECT().PutVehicle(*vehicle).Times(1)

	updateMileageCommand := update_mileage_command.NewUpdateMileageCommand(deviceSerialNumber, mileage)

	updateMileageCommandHandler := update_mileage_command.NewUpdateMileageCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateMileageCommandHandler.Handle(updateMileageCommand)

	assert.Nil(t, err)
}

func TestShouldReturnErrorIfVehicleIsNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	mileage := 50

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, mileageVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOMileageVO(inFleetDate, chassisNumber, deviceSerialNumber, mileage)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)
	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, custom_error.EntityNotFound("Vehicle with Device Serial Number abc not found")).Times(1)

	err = vehicle.SetMileage(mileageVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)


	updateMileageCommand := update_mileage_command.NewUpdateMileageCommand(deviceSerialNumber, mileage)

	updateMileageCommandHandler := update_mileage_command.NewUpdateMileageCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateMileageCommandHandler.Handle(updateMileageCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Vehicle with Device Serial Number abc not found", err.Error())
}

func TestShouldReturnErrorIfMileageIsNegative(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	mileage := 50

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, mileageVO, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOMileageVO(inFleetDate, chassisNumber, deviceSerialNumber, mileage)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)
	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	err = vehicle.SetMileage(mileageVO)

	assert.Nil(t, err)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)


	updateMileageCommand := update_mileage_command.NewUpdateMileageCommand(deviceSerialNumber, -10)

	updateMileageCommandHandler := update_mileage_command.NewUpdateMileageCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateMileageCommandHandler.Handle(updateMileageCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Current Mileage -10, the number should be positive integer", err.Error())
}

func TestShouldReturnErrorIfVehicleIsNotInstalled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"
	deviceSerialNumber := "abc"
	mileage := 50

	inFleetDateVO, chassisNumberVO, deviceSerialNumberVO, _, err := createInFleetDateVOChassisNumberVODeviceSerialNumberVOMileageVO(inFleetDate, chassisNumber, deviceSerialNumber, mileage)

	assert.Nil(t, err)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)
	vehicleReadRepositoryMocked.EXPECT().GetVehicleByDeviceSerialNbr(deviceSerialNumberVO).Return(*vehicle, nil).Times(1)

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)
	
	updateMileageCommand := update_mileage_command.NewUpdateMileageCommand(deviceSerialNumber, mileage)

	updateMileageCommandHandler := update_mileage_command.NewUpdateMileageCommandHandler(vehicleWriteRepositoryMocked, vehicleReadRepositoryMocked)

	err = updateMileageCommandHandler.Handle(updateMileageCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't set the mileage of the vehicle if it hasn't been in fleet before", err.Error())
}