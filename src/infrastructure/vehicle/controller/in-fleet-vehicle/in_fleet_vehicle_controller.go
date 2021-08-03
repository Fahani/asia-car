package in_fleet_vehicle_controller

import (
	"encoding/json"
	"fmt"
	"github.com/fahani/asia-car/src/application/vehicle/command/in-fleet-vehicle"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
	"net/http"
)

type InFleetVehicleController struct {
	inFleetVehicleCommandHandler *in_fleet_vehicle_command.InFleetVehicleCommandHandler
}
type InFleetVehicleUnmarshal struct {
	InFleetDate   string `json:"in_fleet_date"`
	ChassisNumber string `json:"chassis_number"`
}

const jsonSchemaValidation = `{
  "type":"object",
  "required":[
    "in_fleet_date",
    "chassis_number"
  ],
  "properties":{
    "in_fleet_date":{
      "type":"string"
    },
    "chassis_number":{
      "type":"string"
    }
  }
}`

func NewInFleetVehicleController(inFleetVehicleCommandHandler *in_fleet_vehicle_command.InFleetVehicleCommandHandler) InFleetVehicleController {
	return InFleetVehicleController{inFleetVehicleCommandHandler: inFleetVehicleCommandHandler}
}

func readBodyRequest(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)

	return string(body), err
}

func isJsonSchemaValid(body, jsonSchemaValidation string) bool {
	schemaLoader := gojsonschema.NewStringLoader(jsonSchemaValidation)
	documentLoader := gojsonschema.NewStringLoader(body)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false
	}

	return result.Valid()
}

func (ifvc *InFleetVehicleController) InFleet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	body, err := readBodyRequest(r)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(500)
		output := `{"error": {"code": 500, "message": "Error reading body request."}}`
		_, err := w.Write([]byte(output))
		if err != nil {
			log.Print(err.Error())
		}
		return
	}

	if isJsonSchemaValid(body, jsonSchemaValidation) == false {
		w.WriteHeader(400)
		output := fmt.Sprintf("{\"error\": {\"code\": 400, \"message\": \"%s\"}}", "The JSON you have provided in your request does not comply with the schema.")
		_, err := w.Write([]byte(output))
		if err != nil {
			log.Print(err.Error())
		}
		return
	}

	var schema InFleetVehicleUnmarshal
	err = json.Unmarshal([]byte(body), &schema)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(500)
		output := fmt.Sprintf("{\"error\": {\"code\": 500, \"message\": \"%s\"}}", err.Error())
		_, err := w.Write([]byte(output))
		if err != nil {
			log.Print(err.Error())
		}
		return
	}

	command := in_fleet_vehicle_command.NewInFleetVehicleCommand(schema.InFleetDate, schema.ChassisNumber)
	err = ifvc.inFleetVehicleCommandHandler.Handle(command)

	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(500)
		output := fmt.Sprintf("{\"error\": {\"code\": 500, \"message\": \"%s\"}}", err.Error())
		_, err := w.Write([]byte(output))
		if err != nil {
			log.Print(err.Error())
		}
		return
	}
	w.WriteHeader(200)
}
