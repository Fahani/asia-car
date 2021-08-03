package update_mileage_command

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/current-mileage"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
)

type UpdateMileageCommandHandler struct {
	vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository
	vehicleReadRepository  vehicle_read_repostiroy.VehicleReadRepository
}

func NewUpdateMileageCommandHandler(vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository, vehicleReadRepository vehicle_read_repostiroy.VehicleReadRepository) *UpdateMileageCommandHandler {
	return &UpdateMileageCommandHandler{
		vehicleWriteRepository: vehicleWriteRepository,
		vehicleReadRepository:  vehicleReadRepository,
	}
}

func (umch *UpdateMileageCommandHandler) Handle(command UpdateMileageCommand) error {
	deviceSerialNumber := device_serial_number.FromValue(command.GetDeviceSerialNumber())

	vehicle, err := umch.vehicleReadRepository.GetVehicleByDeviceSerialNbr(deviceSerialNumber)

	if err != nil {
		return err
	}

	mileage, err := current_mileage.FromValue(command.GetMileageLevel())

	if err != nil {
		return err
	}

	err = vehicle.SetMileage(mileage)

	if err != nil {
		return err
	}

	umch.vehicleWriteRepository.PutVehicle(vehicle)

	return nil
}
