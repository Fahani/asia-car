package vehicle_aggregate

import (
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/battery-level"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/brand"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/category"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/current-mileage"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/fuel-level"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/licence-plate"
)

type Vehicle struct {
	chassisNbr         chassis_nbr.ChassisNbr
	licensePlate       licence_plate.LicencePlate
	brand              brand.Brand
	category           category.Category
	inFleetDate        in_fleet_date.InFleetDate
	deviceSerialNumber device_serial_number.DeviceSerialNumber
	batteryLevel       battery_level.BatteryLevel
	fuelLevel          fuel_level.FuelLevel
	currentMileage     current_mileage.CurrentMileage
}

func InFleetVehicle(chassisNbr chassis_nbr.ChassisNbr, inFleetDate in_fleet_date.InFleetDate) *Vehicle {
	return &Vehicle{chassisNbr: chassisNbr, inFleetDate: inFleetDate}
}

func (vehicle *Vehicle) InstallVehicle(deviceSerialNumber device_serial_number.DeviceSerialNumber) error {
	if !vehicle.isInFleet() {
		return custom_error.InstallVehicle("We can't install a vehicle if hasn't been in fleet before")
	}

	vehicle.deviceSerialNumber = deviceSerialNumber
	return nil
}

func (vehicle *Vehicle) GetChassisNbr() chassis_nbr.ChassisNbr {
	return vehicle.chassisNbr
}

func (vehicle *Vehicle) GetInFleetDate() in_fleet_date.InFleetDate {
	return vehicle.inFleetDate
}

func (vehicle *Vehicle) GetDeviceSerialNbr() device_serial_number.DeviceSerialNumber {
	return vehicle.deviceSerialNumber
}

func (vehicle *Vehicle) isInFleet() bool {
	if vehicle.chassisNbr == (chassis_nbr.ChassisNbr{}) || vehicle.inFleetDate == (in_fleet_date.InFleetDate{}) {
		return false
	}
	return true
}

func (vehicle *Vehicle) isInstalled() bool {
	if vehicle.deviceSerialNumber == (device_serial_number.DeviceSerialNumber{}) {
		return false
	}
	return true
}

func (vehicle *Vehicle) SetMileage(newMileage current_mileage.CurrentMileage) error {
	if !vehicle.isInstalled() {
		return custom_error.InstallVehicle("We can't set the mileage of the vehicle if it hasn't been in fleet before")
	}
	vehicle.currentMileage = newMileage
	return nil
}

func (vehicle *Vehicle) GetMileage() current_mileage.CurrentMileage {
	return vehicle.currentMileage
}

func (vehicle *Vehicle) IncreaseFuel(fuel fuel_level.FuelLevel) error {
	if !vehicle.isInstalled() {
		return custom_error.InstallVehicle("We can't increase the fuel of the vehicle if it hasn't been in fleet before")
	}

	newFuel, _ := fuel_level.FromValue(vehicle.fuelLevel.GetValue() + fuel.GetValue())

	vehicle.fuelLevel = newFuel
	return nil
}

func (vehicle *Vehicle) DecreaseFuel(fuel fuel_level.FuelLevel) error {
	if !vehicle.isInstalled() {
		return custom_error.InstallVehicle("We can't decrease the fuel level of the vehicle if it hasn't been in fleet before")
	}
	newFuel, err := fuel_level.FromValue(vehicle.fuelLevel.GetValue() - fuel.GetValue())
	if err != nil {
		return err
	}

	vehicle.fuelLevel = newFuel
	return nil
}

func (vehicle *Vehicle) GetFuel() fuel_level.FuelLevel {
	return vehicle.fuelLevel
}

func (vehicle *Vehicle) SetBatteryLevel(newBattery battery_level.BatteryLevel) error {
	if !vehicle.isInstalled() {
		return custom_error.InstallVehicle("We can't set the battery level of the vehicle if it hasn't been in fleet before")
	}
	vehicle.batteryLevel = newBattery
	return nil
}

func (vehicle *Vehicle) GetBatteryLevel() battery_level.BatteryLevel {
	return vehicle.batteryLevel
}

func (vehicle *Vehicle) SetLicencePlate(newLicencePlate licence_plate.LicencePlate) error {
	if !vehicle.isInstalled() {
		return custom_error.InstallVehicle("We can't set the licence plate of the vehicle if it hasn't been in fleet before")
	}
	vehicle.licensePlate = newLicencePlate
	return nil
}

func (vehicle *Vehicle) GetLicencePlate() licence_plate.LicencePlate {
	return vehicle.licensePlate
}

func (vehicle *Vehicle) SetBrand(newBrand brand.Brand) error {
	if !vehicle.isInstalled() {
		return custom_error.InstallVehicle("We can't set the brand of the vehicle if it hasn't been in fleet before")
	}
	vehicle.brand = newBrand
	return nil
}

func (vehicle *Vehicle) GetBrand() brand.Brand {
	return vehicle.brand
}

func (vehicle *Vehicle) SetCategory(newCategory category.Category) error {
	if !vehicle.isInstalled() {
		return custom_error.InstallVehicle("We can't set the category of the vehicle if it hasn't been in fleet before")
	}
	vehicle.category = newCategory
	return nil
}

func (vehicle *Vehicle) GetCategory() category.Category {
	return vehicle.category
}
