package vehicle_write_repostiroy

import "github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"

type VehicleWriteRepository interface {
	PutVehicle(vehicle vehicle_aggregate.Vehicle)
}
