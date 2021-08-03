package in_fleet_vehicle_controller_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fahani/asia-car/src/application/vehicle/command/in-fleet-vehicle"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/write"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/controller/in-fleet-vehicle"
	"github.com/fahani/asia-car/src/infrastructure/vehicle/write-model/repository/read-write"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func setup(vehicleWriteRepository vehicle_write_repostiroy.VehicleWriteRepository, port, path string) *http.Server {
	srv := http.Server{Addr: fmt.Sprintf(":%s", port)}

	handler := in_fleet_vehicle_command.NewInFleetVehicleCommandHandler(vehicleWriteRepository)
	controller := in_fleet_vehicle_controller.NewInFleetVehicleController(handler)

	http.HandleFunc(path, controller.InFleet)
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

func TestVehicleIsInFleet(t *testing.T) {
	port := "9001"
	path := "/vehicles/in-fleet"
	vehicleWriteRepository := in_memory_vehicle_read_write_repository.NewInMemoryVehicleRepository()
	serv := setup(vehicleWriteRepository, port, path)
	body := `{"in_fleet_date": "2014-11-12T11:45:26.371Z", "chassis_number": "01234567890123456"}`
	resp, err := http.Post(fmt.Sprintf("http://localhost:%s%s", port, path), "application/json", bytes.NewBuffer([]byte(body)))
	assert.Nil(t, err)

	responseBody, err := getBodyString(resp)

	assert.Equal(t, "", responseBody)

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	err = teardown(serv)
	assert.Nil(t, err)

	chassisNumberVO, _ := chassis_nbr.FromValue("01234567890123456")
	vehicle, err := vehicleWriteRepository.GetVehicleByChassisNbr(chassisNumberVO)

	assert.Nil(t, err)

	assert.Equal(t, "01234567890123456", vehicle.GetChassisNbr().GetValue())
	assert.Equal(t, "2014-11-12 11:45:26.371 +0000 UTC", vehicle.GetInFleetDate().GetValue().String())
}


func TestShouldGetErrorWhenBodyDoesNotComplyWithSchema(t *testing.T) {
	port := "9002"
	path := "/vehicles/in-fleet2"
	vehicleWriteRepository := in_memory_vehicle_read_write_repository.NewInMemoryVehicleRepository()
	serv := setup(vehicleWriteRepository, port, path)
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