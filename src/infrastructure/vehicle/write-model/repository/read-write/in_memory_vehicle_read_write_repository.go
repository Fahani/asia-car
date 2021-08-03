package in_memory_vehicle_read_write_repository

import (
	"fmt"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"sync"
)

type InMemoryVehicleRepository struct {
	vehiclesByChassisNbr      map[string]vehicle_aggregate.Vehicle
	vehiclesByDeviceSerialNbr map[string]vehicle_aggregate.Vehicle
}

func NewInMemoryVehicleRepository() *InMemoryVehicleRepository {
	return &InMemoryVehicleRepository{
		vehiclesByChassisNbr:      make(map[string]vehicle_aggregate.Vehicle),
		vehiclesByDeviceSerialNbr: make(map[string]vehicle_aggregate.Vehicle),
	}
}

func (imvr *InMemoryVehicleRepository) PutVehicle(vehicle vehicle_aggregate.Vehicle) {
	// For the shake of using Locks. Note that this could only make sense when using one instance.
	var mu sync.Mutex
	mu.Lock()
	if vehicle.GetChassisNbr() != (chassis_nbr.ChassisNbr{}) {
		imvr.vehiclesByChassisNbr[vehicle.GetChassisNbr().GetValue()] = vehicle
	}
	if vehicle.GetDeviceSerialNbr() != (device_serial_number.DeviceSerialNumber{}) {
		imvr.vehiclesByDeviceSerialNbr[vehicle.GetDeviceSerialNbr().GetValue()] = vehicle
	}
	mu.Unlock()
}

func (imvr *InMemoryVehicleRepository) GetVehicleByChassisNbr(chassisNbr chassis_nbr.ChassisNbr) (vehicle_aggregate.Vehicle, error) {
	vehicle, ok := imvr.vehiclesByChassisNbr[chassisNbr.GetValue()]
	if !ok {
		return vehicle_aggregate.Vehicle{}, custom_error.EntityNotFound(fmt.Sprintf("Vehicle with Chassis Number %s not found", chassisNbr.GetValue()))
	}
	return vehicle, nil
}

func (imvr *InMemoryVehicleRepository) GetVehicleByDeviceSerialNbr(deviceSerialNumber device_serial_number.DeviceSerialNumber) (vehicle_aggregate.Vehicle, error) {
	vehicle, ok := imvr.vehiclesByDeviceSerialNbr[deviceSerialNumber.GetValue()]
	if !ok {
		return vehicle_aggregate.Vehicle{}, custom_error.EntityNotFound(fmt.Sprintf("Vehicle with Device Serial Number %s not found", deviceSerialNumber.GetValue()))
	}
	return vehicle, nil
}
