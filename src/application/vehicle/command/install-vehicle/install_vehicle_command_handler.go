package install_vehicle_command

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
)

type InstallVehicleCommandHandler struct {
	vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository
	vehicleReadRepository  vehicle_read_repostiroy.VehicleReadRepository
}

func NewInstallVehicleCommandHandler(vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository, vehicleReadRepository vehicle_read_repostiroy.VehicleReadRepository) *InstallVehicleCommandHandler {
	return &InstallVehicleCommandHandler{
		vehicleWriteRepository: vehicleWriteRepository,
		vehicleReadRepository:  vehicleReadRepository,
	}
}

func (ifvch *InstallVehicleCommandHandler) Handle(command InstallVehicleCommand) error {
	deviceSerialNumber := device_serial_number.FromValue(command.GetDeviceSerialNumber())

	cn, err := chassis_nbr.FromValue(command.GetChassisNumber())

	if err != nil {
		return err
	}

	v, err := ifvch.vehicleReadRepository.GetVehicleByChassisNbr(cn)

	if err != nil {
		return err
	}

	err = v.InstallVehicle(deviceSerialNumber)

	if err != nil {
		return err
	}

	ifvch.vehicleWriteRepository.PutVehicle(v)

	return nil
}
