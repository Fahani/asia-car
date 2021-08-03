package get_telemetries_by_chassis_number_query_test

import (
	"github.com/fahani/asia-car/src/application/vehicle/query/get-telemetries-by-chassis-number"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read/mock"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/battery-level"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/brand"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/category"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/current-mileage"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/device-serial-number"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/fuel-level"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/licence-plate"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createVehicleVOs(chassisNumber, licencePlate, brandString, categoryString, inFleetDate, deviceSerialNumber string, battery, fuel, mileage int) (chassis_nbr.ChassisNbr, licence_plate.LicencePlate, brand.Brand, category.Category, in_fleet_date.InFleetDate, device_serial_number.DeviceSerialNumber, battery_level.BatteryLevel, fuel_level.FuelLevel, current_mileage.CurrentMileage, error) {
	chassisNumberVO, err := chassis_nbr.FromValue(chassisNumber)

	if err != nil {
		return chassis_nbr.ChassisNbr{}, licence_plate.LicencePlate{}, brand.Brand{}, category.Category{}, in_fleet_date.InFleetDate{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, fuel_level.FuelLevel{}, current_mileage.CurrentMileage{}, err
	}

	licencePlateVO, err := licence_plate.FromValue(licencePlate)

	if err != nil {
		return chassis_nbr.ChassisNbr{}, licence_plate.LicencePlate{}, brand.Brand{}, category.Category{}, in_fleet_date.InFleetDate{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, fuel_level.FuelLevel{}, current_mileage.CurrentMileage{}, err
	}

	brandVO, err := brand.FromValue(brandString)

	if err != nil {
		return chassis_nbr.ChassisNbr{}, licence_plate.LicencePlate{}, brand.Brand{}, category.Category{}, in_fleet_date.InFleetDate{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, fuel_level.FuelLevel{}, current_mileage.CurrentMileage{}, err
	}

	categoryVO, err := category.FromValue(categoryString)

	if err != nil {
		return chassis_nbr.ChassisNbr{}, licence_plate.LicencePlate{}, brand.Brand{}, category.Category{}, in_fleet_date.InFleetDate{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, fuel_level.FuelLevel{}, current_mileage.CurrentMileage{}, err
	}

	t, err := time.Parse(time.RFC3339, inFleetDate)

	if err != nil {
		return chassis_nbr.ChassisNbr{}, licence_plate.LicencePlate{}, brand.Brand{}, category.Category{}, in_fleet_date.InFleetDate{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, fuel_level.FuelLevel{}, current_mileage.CurrentMileage{}, err
	}

	inFleetDateVO := in_fleet_date.FromValue(t)

	deviceSerialNumberVO := device_serial_number.FromValue(deviceSerialNumber)

	batteryVO, err := battery_level.FromValue(battery)

	if err != nil {
		return chassis_nbr.ChassisNbr{}, licence_plate.LicencePlate{}, brand.Brand{}, category.Category{}, in_fleet_date.InFleetDate{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, fuel_level.FuelLevel{}, current_mileage.CurrentMileage{}, err
	}

	fuelVO, err := fuel_level.FromValue(fuel)

	if err != nil {
		return chassis_nbr.ChassisNbr{}, licence_plate.LicencePlate{}, brand.Brand{}, category.Category{}, in_fleet_date.InFleetDate{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, fuel_level.FuelLevel{}, current_mileage.CurrentMileage{}, err
	}

	mileageVO, err := current_mileage.FromValue(mileage)

	if err != nil {
		return chassis_nbr.ChassisNbr{}, licence_plate.LicencePlate{}, brand.Brand{}, category.Category{}, in_fleet_date.InFleetDate{}, device_serial_number.DeviceSerialNumber{}, battery_level.BatteryLevel{}, fuel_level.FuelLevel{}, current_mileage.CurrentMileage{}, err
	}

	return chassisNumberVO, licencePlateVO, brandVO, categoryVO, inFleetDateVO, deviceSerialNumberVO, batteryVO, fuelVO, mileageVO, nil
}

func TestShouldReturnTelemetriesResultCorrectly(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)

	chassisNumber := "01234567890123456"
	licencePlate := "abc"
	brandString := "seat"
	categoryString := "MBMR"
	inFleetDate := "2014-11-12T11:45:26.371Z"
	deviceSerialNumber := "abc"
	battery := 50
	fuel := 50
	mileage := 50


	chassisNumberVO, licencePlateVO, brandVO, categoryVO, inFleetDateVO, deviceSerialNumberVO, batteryVO, fuelVO, mileageVO, err := createVehicleVOs(chassisNumber,licencePlate, brandString, categoryString, inFleetDate, deviceSerialNumber, battery, fuel, mileage)

	assert.Nil(t, err)


	getTelemetriesBySerialNumberQueryResponse := get_telemetries_by_chassis_number_query.GetTelemetriesByChassisNumberQueryResponse{
		ChassisNumber:      chassisNumber,
		LicensePlate:       licencePlate,
		Brand:              brandString,
		Category:           categoryString,
		InFleetDate:        inFleetDateVO.GetValue(),
		DeviceSerialNumber: deviceSerialNumber,
		BatteryLevel:       battery,
		FuelLevel:          fuel,
		CurrentMileage:     mileage,
	}

	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDateVO)

	err = vehicle.InstallVehicle(deviceSerialNumberVO)

	assert.Nil(t, err)

	err = vehicle.SetLicencePlate(licencePlateVO)

	assert.Nil(t, err)

	err = vehicle.SetBrand(brandVO)

	assert.Nil(t, err)

	err = vehicle.SetCategory(categoryVO)

	assert.Nil(t, err)

	err = vehicle.SetBatteryLevel(batteryVO)

	assert.Nil(t, err)

	err = vehicle.IncreaseFuel(fuelVO)

	assert.Nil(t, err)

	err = vehicle.SetMileage(mileageVO)

	assert.Nil(t, err)

	vehicleReadRepositoryMocked.EXPECT().GetVehicleByChassisNbr(chassisNumberVO).Return(*vehicle, nil).Times(1)

	handler := get_telemetries_by_chassis_number_query.NewGetTelemetriesByChassisNumberQueryHandler(vehicleReadRepositoryMocked)

	query := get_telemetries_by_chassis_number_query.NewGetTelemetriesByChassisNumberQuery(chassisNumber)

	response, err := handler.Handle(query)

	assert.Nil(t, err)

	assert.Equal(t, getTelemetriesBySerialNumberQueryResponse.ChassisNumber, response.ChassisNumber)
	assert.Equal(t, getTelemetriesBySerialNumberQueryResponse.LicensePlate, response.LicensePlate)
	assert.Equal(t, getTelemetriesBySerialNumberQueryResponse.Brand, response.Brand)
	assert.Equal(t, getTelemetriesBySerialNumberQueryResponse.Category, response.Category)
	assert.Equal(t, getTelemetriesBySerialNumberQueryResponse.InFleetDate, response.InFleetDate)
	assert.Equal(t, getTelemetriesBySerialNumberQueryResponse.DeviceSerialNumber, response.DeviceSerialNumber)
	assert.Equal(t, getTelemetriesBySerialNumberQueryResponse.BatteryLevel, response.BatteryLevel)
	assert.Equal(t, getTelemetriesBySerialNumberQueryResponse.FuelLevel, response.FuelLevel)
	assert.Equal(t, getTelemetriesBySerialNumberQueryResponse.CurrentMileage, response.CurrentMileage)
}

func TestShouldReturnErrorIfVehicleNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)

	chassisNumber := "01234567890123456"
	licencePlate := "abc"
	brandString := "seat"
	categoryString := "MBMR"
	inFleetDate := "2014-11-12T11:45:26.371Z"
	deviceSerialNumber := "abc"
	battery := 50
	fuel := 50
	mileage := 50

	chassisNumberVO, _, _, _, _, _, _, _, _, err := createVehicleVOs(chassisNumber,licencePlate, brandString, categoryString, inFleetDate, deviceSerialNumber, battery, fuel, mileage)

	assert.Nil(t, err)

	vehicleReadRepositoryMocked.EXPECT().GetVehicleByChassisNbr(chassisNumberVO).Return(vehicle_aggregate.Vehicle{}, custom_error.EntityNotFound("Vehicle with Chassis Number 01234567890123456 not found")).Times(1)

	handler := get_telemetries_by_chassis_number_query.NewGetTelemetriesByChassisNumberQueryHandler(vehicleReadRepositoryMocked)

	query := get_telemetries_by_chassis_number_query.NewGetTelemetriesByChassisNumberQuery(chassisNumber)

	_, err = handler.Handle(query)

	assert.NotNil(t, err)
	assert.Equal(t, "Vehicle with Chassis Number 01234567890123456 not found", err.Error())
}

func TestShouldReturnErrorIfChassisIsNotValid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	vehicleReadRepositoryMocked := mock_vehicle_read_repostiroy.NewMockVehicleReadRepository(mockCtrl)

	chassisNumber := "0123456789012"

	handler := get_telemetries_by_chassis_number_query.NewGetTelemetriesByChassisNumberQueryHandler(vehicleReadRepositoryMocked)

	query := get_telemetries_by_chassis_number_query.NewGetTelemetriesByChassisNumberQuery(chassisNumber)

	_, err := handler.Handle(query)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid Chassis Number 0123456789012, the number should be 17 digits alphanumeric", err.Error())
}