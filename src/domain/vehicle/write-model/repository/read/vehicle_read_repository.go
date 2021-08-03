package vehicle_read_repostiroy

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
)

type VehicleReadRepository interface {
	GetVehicleByChassisNbr(chassisNbr chassis_nbr.ChassisNbr) (vehicle_aggregate.Vehicle, error)
	GetVehicleByDeviceSerialNbr(deviceSerialNumber device_serial_number.DeviceSerialNumber) (vehicle_aggregate.Vehicle, error)
}
