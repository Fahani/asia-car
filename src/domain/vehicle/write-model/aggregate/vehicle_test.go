package vehicle_aggregate_test

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/battery-level"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/brand"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/category"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/current-mileage"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/fuel-level"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/licence-plate"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestShouldInFleetAVehicleCorrectly(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	assert.Equal(t, expectedChassisNbr, vehicle.GetChassisNbr().GetValue())
	assert.Equal(t, nowInFleetTime, vehicle.GetInFleetDate().GetValue())
}

func TestShouldFailToInstallIfNotFleetBefore(t *testing.T) {
	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	cn, _ := chassis_nbr.FromValue("01234567890123456")
	vs := vehicle_aggregate.InFleetVehicle(cn, in_fleet_date.InFleetDate{})
	err := vs.InstallVehicle(deviceSerialNbr)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't install a vehicle if hasn't been in fleet before", err.Error())

}

func TestShouldInstallVehicleCorrectly(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	_ = vehicle.InstallVehicle(deviceSerialNbr)

	assert.Equal(t, expectedDeviceNbr, vehicle.GetDeviceSerialNbr().GetValue())
}

func TestShouldSetAndGetTheMileageCorrectly(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	_ = vehicle.InstallVehicle(deviceSerialNbr)

	mileage, _ := current_mileage.FromValue(100)

	err := vehicle.SetMileage(mileage)

	assert.Nil(t, err)
	assert.Equal(t, mileage.GetValue(), vehicle.GetMileage().GetValue())
}

func TestShouldNotSetTheMileageIfVehicleIsNotInstalled(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	mileage, _ := current_mileage.FromValue(100)

	err := vehicle.SetMileage(mileage)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't set the mileage of the vehicle if it hasn't been in fleet before", err.Error())
}

func TestShouldIncreaseFuelCorrectly(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	_ = vehicle.InstallVehicle(deviceSerialNbr)

	fuel, _ := fuel_level.FromValue(100)

	err := vehicle.IncreaseFuel(fuel)

	assert.Nil(t, err)
	assert.Equal(t, fuel.GetValue(), vehicle.GetFuel().GetValue())
}

func TestShouldReturnErrorWhenIncreaseFuelAndVehicleIsNotInstalled(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	fuel, _ := fuel_level.FromValue(100)

	err := vehicle.IncreaseFuel(fuel)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't increase the fuel of the vehicle if it hasn't been in fleet before", err.Error())
}

func TestShouldDecreaseFuelCorrectly(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	_ = vehicle.InstallVehicle(deviceSerialNbr)

	fuel, _ := fuel_level.FromValue(100)

	err := vehicle.IncreaseFuel(fuel)

	assert.Nil(t, err)
	assert.Equal(t, fuel.GetValue(), vehicle.GetFuel().GetValue())

	err = vehicle.DecreaseFuel(fuel)

	assert.Nil(t, err)
	assert.Equal(t, 0, vehicle.GetFuel().GetValue())
}

func TestShouldReturnErrorWhenDecreasingFuelAndVehicleNotInstalled(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)


	fuel, _ := fuel_level.FromValue(100)

	err := vehicle.DecreaseFuel(fuel)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't decrease the fuel level of the vehicle if it hasn't been in fleet before", err.Error())
}

func TestShouldReturnErrorIfDecreasingFuelIsBelowZero(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	_ = vehicle.InstallVehicle(deviceSerialNbr)

	fuel, _ := fuel_level.FromValue(100)

	err := vehicle.DecreaseFuel(fuel)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Fuel Level -100, the number should be positive integer", err.Error())
}

func TestShouldSetAndGetTheBatteryLevelCorrectly(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	_ = vehicle.InstallVehicle(deviceSerialNbr)

	battery, _ := battery_level.FromValue(100)

	err := vehicle.SetBatteryLevel(battery)

	assert.Nil(t, err)
	assert.Equal(t, battery.GetValue(), vehicle.GetBatteryLevel().GetValue())
}

func TestShouldReturnErrorSettingBatteryLevelAndVehicleNotInstalled(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	battery, _ := battery_level.FromValue(100)

	err := vehicle.SetBatteryLevel(battery)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't set the battery level of the vehicle if it hasn't been in fleet before", err.Error())
}

func TestShouldSetAndGetTheLicencePlateCorrectly(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	_ = vehicle.InstallVehicle(deviceSerialNbr)

	licencePlate, _ := licence_plate.FromValue("AA")

	err := vehicle.SetLicencePlate(licencePlate)

	assert.Nil(t, err)
	assert.Equal(t, licencePlate.GetValue(), vehicle.GetLicencePlate().GetValue())
}

func TestShouldReturnErrorSettingTheLicencePlateIfVehicleIsNotInstalled(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	licencePlate, _ := licence_plate.FromValue("AA")

	err := vehicle.SetLicencePlate(licencePlate)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't set the licence plate of the vehicle if it hasn't been in fleet before", err.Error())
}

func TestShouldSetAndGetTheBrandCorrectly(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	_ = vehicle.InstallVehicle(deviceSerialNbr)

	newBrand, _ := brand.FromValue("Seat")

	err := vehicle.SetBrand(newBrand)

	assert.Nil(t, err)
	assert.Equal(t, newBrand.GetValue(), vehicle.GetBrand().GetValue())
}

func TestShouldReturnAnErrorWhenSettingTheBrandAndVehicleIsNotInstalled(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	newBrand, _ := brand.FromValue("Seat")

	err := vehicle.SetBrand(newBrand)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't set the brand of the vehicle if it hasn't been in fleet before", err.Error())
}

func TestShouldSetAndGetTheCategoryCorrectly(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	expectedDeviceNbr := "abc"
	deviceSerialNbr := device_serial_number.FromValue(expectedDeviceNbr)

	_ = vehicle.InstallVehicle(deviceSerialNbr)

	newCategory, _ := category.FromValue("XKDX")

	err := vehicle.SetCategory(newCategory)

	assert.Nil(t, err)
	assert.Equal(t, newCategory.GetValue(), vehicle.GetCategory().GetValue())
}

func TestShouldReturnErrorWhenSettingTheCategoryAndVehicleIsNotInstalled(t *testing.T) {
	expectedChassisNbr := "01234567890123456"
	chassisNbr, _ := chassis_nbr.FromValue(expectedChassisNbr)

	nowInFleetTime := time.Now()
	inFleetDate := in_fleet_date.FromValue(nowInFleetTime)

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNbr, inFleetDate)

	newCategory, _ := category.FromValue("XKDX")

	err := vehicle.SetCategory(newCategory)

	assert.NotNil(t, err)
	assert.Equal(t, "We can't set the category of the vehicle if it hasn't been in fleet before", err.Error())
}