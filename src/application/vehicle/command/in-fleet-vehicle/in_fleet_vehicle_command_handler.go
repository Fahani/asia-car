package in_fleet_vehicle_command

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"time"
)

type InFleetVehicleCommandHandler struct {
	vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository
}

func NewInFleetVehicleCommandHandler(vehicleRepository vehicle_write_repostiroy.VehicleWriteRepository) *InFleetVehicleCommandHandler {
	return &InFleetVehicleCommandHandler{vehicleWriteRepository: vehicleRepository}
}

func (ifvch *InFleetVehicleCommandHandler) Handle(command InFleetVehicleCommand) error {
	chassisNumber, err := chassis_nbr.FromValue(command.GetChassisNumber())

	if err != nil {
		return err
	}

	t, err := time.Parse(time.RFC3339, command.GetInFleetDate())

	if err != nil {
		return err
	}

	inFleetDate := in_fleet_date.FromValue(t)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumber, inFleetDate)
	ifvch.vehicleWriteRepository.PutVehicle(*vehicle)

	return nil
}
