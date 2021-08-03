package update_battery_command

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/battery-level"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
)

type UpdateBatteryLevelCommandHandler struct {
	vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository
	vehicleReadRepository  vehicle_read_repostiroy.VehicleReadRepository
}

func NewUpdateBatteryLevelCommandHandler(vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository, vehicleReadRepository vehicle_read_repostiroy.VehicleReadRepository) *UpdateBatteryLevelCommandHandler {
	return &UpdateBatteryLevelCommandHandler{
		vehicleWriteRepository: vehicleWriteRepository,
		vehicleReadRepository:  vehicleReadRepository,
	}
}

func (ublch *UpdateBatteryLevelCommandHandler) Handle(command UpdateBatteryCommand) error {
	deviceSerialNumber := device_serial_number.FromValue(command.GetDeviceSerialNumber())

	vehicle, err := ublch.vehicleReadRepository.GetVehicleByDeviceSerialNbr(deviceSerialNumber)

	if err != nil {
		return err
	}

	batteryLevel, err := battery_level.FromValue(command.GetBatteryLevel())

	if err != nil {
		return err
	}

	err = vehicle.SetBatteryLevel(batteryLevel)

	if err != nil {
		return err
	}

	ublch.vehicleWriteRepository.PutVehicle(vehicle)

	return nil
}
