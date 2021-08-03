// +build wireinject

package main

import (
	"github.com/fahani/asia-car/src/application/vehicle/command/in-fleet-vehicle"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/write-model/repository/read-write"
	"github.com/google/wire"
)

func InitializeInFleetCommandHandler() *in_fleet_vehicle_command.InFleetVehicleCommandHandler {
	wire.Build(
		wire.Bind(new(vehicle_write_repostiroy.VehicleWriteRepository), new(*in_memory_vehicle_read_write_repository.InMemoryVehicleRepository)),
		in_memory_vehicle_read_write_repository.NewInMemoryVehicleRepository,
		in_fleet_vehicle_command.NewInFleetVehicleCommandHandler)

	return &in_fleet_vehicle_command.InFleetVehicleCommandHandler{}
}
