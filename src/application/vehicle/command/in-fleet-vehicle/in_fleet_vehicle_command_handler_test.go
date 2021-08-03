package in_fleet_vehicle_command_test

import (
	"github.com/fahani/asia-car/src/application/vehicle/command/in-fleet-vehicle"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createInFleetDateAndChassisNumberVO(inFleetDate, chassisNumber string) (in_fleet_date.InFleetDate, chassis_nbr.ChassisNbr, error) {
	chassisNumberVO, err := chassis_nbr.FromValue(chassisNumber)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, err
	}

	t, err := time.Parse(time.RFC3339, inFleetDate)

	if err != nil {
		return in_fleet_date.InFleetDate{}, chassis_nbr.ChassisNbr{}, err
	}

	inFleetDateVO := in_fleet_date.FromValue(t)

	return inFleetDateVO, chassisNumberVO, nil
}

func TestShouldInFleetVehicleCorrectly(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "01234567890123456"

	inFleetDateVO, chassisNumberVo, err := createInFleetDateAndChassisNumberVO(inFleetDate, chassisNumber)

	assert.Nil(t, err)

	vehicleInFleet := vehicle_aggregate.InFleetVehicle(chassisNumberVo, inFleetDateVO)
	vehicleWriteRepositoryMocked.EXPECT().PutVehicle(*vehicleInFleet).Times(1)

	handler := in_fleet_vehicle_command.NewInFleetVehicleCommandHandler(vehicleWriteRepositoryMocked)

	vehicleInFleetCommand := in_fleet_vehicle_command.NewInFleetVehicleCommand(inFleetDate, chassisNumber)

	err = handler.Handle(vehicleInFleetCommand)

	assert.Nil(t, err)
}

func TestShouldReturnErrorIfChassisNumberIsNotValid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	inFleetDate := "2014-11-12T11:45:26.371Z"
	chassisNumber := "012345678901234"

	handler := in_fleet_vehicle_command.NewInFleetVehicleCommandHandler(vehicleWriteRepositoryMocked)

	vehicleInFleetCommand := in_fleet_vehicle_command.NewInFleetVehicleCommand(inFleetDate, chassisNumber)

	err := handler.Handle(vehicleInFleetCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Chassis Number 012345678901234, the number should be 17 digits alphanumeric", err.Error())
}

func TestShouldReturnErrorIfInFleetDateIsNotValid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	vehicleWriteRepositoryMocked := mock_vehicle_write_repostiroy.NewMockVehicleWriteRepository(mockCtrl)

	inFleetDate := "Invalid"
	chassisNumber := "01234567890123456"

	handler := in_fleet_vehicle_command.NewInFleetVehicleCommandHandler(vehicleWriteRepositoryMocked)

	vehicleInFleetCommand := in_fleet_vehicle_command.NewInFleetVehicleCommand(inFleetDate, chassisNumber)

	err := handler.Handle(vehicleInFleetCommand)

	assert.NotNil(t, err)
	assert.Equal(t, "parsing time \"Invalid\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"Invalid\" as \"2006\"", err.Error())
}
