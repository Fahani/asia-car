package update_fuel_command

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/fuel-level"
)

const updateTypeIncrement = "increment"

type UpdateCommandHandler struct {
	vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository
	vehicleReadRepository  vehicle_read_repostiroy.VehicleReadRepository
}

func NewUpdateFuelCommandHandler(vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository, vehicleReadRepository vehicle_read_repostiroy.VehicleReadRepository) *UpdateCommandHandler {
	return &UpdateCommandHandler{
		vehicleWriteRepository: vehicleWriteRepository,
		vehicleReadRepository:  vehicleReadRepository,
	}
}

func (uch *UpdateCommandHandler) Handle(command UpdateFuelCommand) error {
	deviceSerialNumber := device_serial_number.FromValue(command.GetDeviceSerialNumber())

	vehicle, err := uch.vehicleReadRepository.GetVehicleByDeviceSerialNbr(deviceSerialNumber)

	if err != nil {
		return err
	}

	fuel, err := fuel_level.FromValue(command.GetFuel())

	if err != nil {
		return err
	}

	if command.GetUpdateType() == updateTypeIncrement {
		err = vehicle.IncreaseFuel(fuel)
	} else {
		err = vehicle.DecreaseFuel(fuel)
	}

	if err != nil {
		return err
	}

	uch.vehicleWriteRepository.PutVehicle(vehicle)

	return nil
}