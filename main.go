package main

import (
	"fmt"
	"github.com/fahani/asia-car/src/application/vehicle/command/in-fleet-vehicle"
	"github.com/fahani/asia-car/src/application/vehicle/command/install-vehicle"
	"github.com/fahani/asia-car/src/application/vehicle/command/update-battery"
	"github.com/fahani/asia-car/src/application/vehicle/command/update-fuel"
	"github.com/fahani/asia-car/src/application/vehicle/command/update-mileage"
	"github.com/fahani/asia-car/src/application/vehicle/query/get-telemetries-by-chassis-number"
	"github.com/fahani/asia-car/src/infrastructure/status/controller"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/controller/in-fleet-vehicle"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/controller/install-vehicle"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/controller/update-battery"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/controller/update-fuel"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/controller/update-mileage"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/controller/vehicle-details"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/write-model/repository/read-write"
	"net/http"
)

func initStatus()  {
	http.HandleFunc("/status", status_controller.Status)
}

func initInFeet(vehicleRepository *in_memory_vehicle_read_write_repository.InMemoryVehicleRepository)  {
	inFleetHandler := in_fleet_vehicle_command.NewInFleetVehicleCommandHandler(vehicleRepository)
	inFleetController := in_fleet_vehicle_controller.NewInFleetVehicleController(inFleetHandler)
	http.HandleFunc("/vehicles/in-fleet", inFleetController.InFleet)
}

func initInstall(vehicleRepository *in_memory_vehicle_read_write_repository.InMemoryVehicleRepository)  {
	installHandler := install_vehicle_command.NewInstallVehicleCommandHandler(vehicleRepository, vehicleRepository)
	installController := install_vehicle_controller.NewInstallVehicleController(installHandler)
	http.HandleFunc("/vehicles/install", installController.Install)
}

func initUpdateBattery(vehicleRepository *in_memory_vehicle_read_write_repository.InMemoryVehicleRepository)  {
	updateBatteryHandler := update_battery_command.NewUpdateBatteryLevelCommandHandler(vehicleRepository, vehicleRepository)
	updateBatteryController := update_battery_controller.NewUpdateBatteryController(updateBatteryHandler)
	http.HandleFunc("/vehicles/update-battery", updateBatteryController.UpdateBattery)
}

func initUpdateFuel(vehicleRepository *in_memory_vehicle_read_write_repository.InMemoryVehicleRepository)  {
	updateFuelHandler := update_fuel_command.NewUpdateFuelCommandHandler(vehicleRepository, vehicleRepository)
	updateFuelController := update_fuel_controller.NewUpdateFuelController(updateFuelHandler)
	http.HandleFunc("/vehicles/update-fuel", updateFuelController.UpdateFuel)
}

func initUpdateMileage(vehicleRepository *in_memory_vehicle_read_write_repository.InMemoryVehicleRepository)  {
	updateMileageHandler := update_mileage_command.NewUpdateMileageCommandHandler(vehicleRepository, vehicleRepository)
	updateMileageController := update_mileage_controller.NewUpdateMileageController(updateMileageHandler)
	http.HandleFunc("/vehicles/update-mileage", updateMileageController.UpdateMileage)
}

func initVehicleDetails(vehicleRepository *in_memory_vehicle_read_write_repository.InMemoryVehicleRepository)  {
	vehicleDetailsHandler := get_telemetries_by_chassis_number_query.NewGetTelemetriesByChassisNumberQueryHandler(vehicleRepository)
	vehicleDetailsController := vehicle_details_controller.NewVehicleDetailsController(vehicleDetailsHandler)
	http.HandleFunc("/vehicles/details", vehicleDetailsController.VehicleDetails)
}

func main()  {
	vehicleRepository := in_memory_vehicle_read_write_repository.NewInMemoryVehicleRepository()

	initStatus()
	initInFeet(vehicleRepository)
	initInstall(vehicleRepository)
	initUpdateBattery(vehicleRepository)
	initUpdateFuel(vehicleRepository)
	initUpdateMileage(vehicleRepository)
	initVehicleDetails(vehicleRepository)

	err := http.ListenAndServe(":9000", nil)
	if nil != err {
		fmt.Println("Error handling requests: " + err.Error())
		return
	}
}
