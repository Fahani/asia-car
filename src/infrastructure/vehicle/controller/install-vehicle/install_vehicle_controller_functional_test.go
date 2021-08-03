package install_vehicle_controller_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fahani/asia-car/src/application/vehicle/command/install-vehicle"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/aggregate"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/in-feet-date"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/controller/install-vehicle"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/write-model/repository/read-write"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func setup(vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository, vehicleReadRepository vehicle_read_repostiroy.VehicleReadRepository, port, path string) *http.Server {
	srv := http.Server{Addr: fmt.Sprintf(":%s", port)}

	handler := install_vehicle_command.NewInstallVehicleCommandHandler(vehicleWriteRepository, vehicleReadRepository)
	controller := install_vehicle_controller.NewInstallVehicleController(handler)

	http.HandleFunc(path, controller.Install)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	time.Sleep(500 * time.Millisecond)
	return &srv
}

func teardown(s *http.Server) error {
	err := s.Close()
	if err != nil {
		return err
	}
	err = s.Shutdown(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func getBodyString(resp *http.Response) (string, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func TestVehicleIsInstalled(t *testing.T) {
	port := "9003"
	path := "/vehicles/install"
	vehicleRepository := in_memory_vehicle_read_write_repository.NewInMemoryVehicleRepository()
	chassisNumberVO, _ := chassis_nbr.FromValue("01234567890123456")
	ti, err := time.Parse(time.RFC3339, "2014-11-12T11:45:26.371Z")
	inFleetDate := in_fleet_date.FromValue(ti)
	vehicle := vehicle_aggregate.InFleetVehicle(chassisNumberVO, inFleetDate)
	vehicleRepository.PutVehicle(*vehicle)
	serv := setup(vehicleRepository, vehicleRepository, port, path)
	body := `{"device_serial_number": "abc", "chassis_number": "01234567890123456"}`
	resp, err := http.Post(fmt.Sprintf("http://localhost:%s%s", port, path), "application/json", bytes.NewBuffer([]byte(body)))
	assert.Nil(t, err)

	responseBody, err := getBodyString(resp)

	assert.Equal(t, "", responseBody)

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	err = teardown(serv)
	assert.Nil(t, err)


	v, err := vehicleRepository.GetVehicleByChassisNbr(chassisNumberVO)

	assert.Nil(t, err)

	assert.Equal(t, "abc", v.GetDeviceSerialNbr().GetValue())
	assert.Equal(t, "2014-11-12 11:45:26.371 +0000 UTC", v.GetInFleetDate().GetValue().String())
	assert.Equal(t, "01234567890123456", v.GetChassisNbr().GetValue())
}

func TestShouldGetErrorWhenBodyDoesNotComplyWithSchema(t *testing.T) {
	port := "9004"
	path := "/vehicles/install2"
	vehicleRepository := in_memory_vehicle_read_write_repository.NewInMemoryVehicleRepository()
	serv := setup(vehicleRepository,vehicleRepository, port, path)
	body := `{"fake": "2014-11-12T11:45:26.371Z", "chassis_number": "01234567890123456"}`
	resp, err := http.Post(fmt.Sprintf("http://localhost:%s%s", port, path), "application/json", bytes.NewBuffer([]byte(body)))
	assert.Nil(t, err)

	responseBody, err := getBodyString(resp)

	assert.Equal(t, "{\"error\": {\"code\": 400, \"message\": \"The JSON you have provided in your request does not comply with the schema.\"}}", responseBody)

	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode)

	err = teardown(serv)
	assert.Nil(t, err)
}
