package install_vehicle_controller

import (
	"encoding/json"
	"fmt"
	"github.com/fahani/asia-car/src/application/vehicle/command/install-vehicle"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
	"net/http"
)

type InstallVehicleController struct {
	installVehicleCommandHandler *install_vehicle_command.InstallVehicleCommandHandler
}
type InstallVehicleUnmarshal struct {
	DeviceSerialNumber string `json:"device_serial_number"`
	ChassisNumber      string `json:"chassis_number"`
}

const jsonSchemaValidation = `{
  "type":"object",
  "required":[
    "device_serial_number",
	"chassis_number"
  ],
  "properties":{
    "device_serial_number":{
      "type":"string"
    },
	"chassis_number":{
      "type":"string"
    }
  }
}`

func NewInstallVehicleController(inFleetVehicleCommandHandler *install_vehicle_command.InstallVehicleCommandHandler) InstallVehicleController {
	return InstallVehicleController{installVehicleCommandHandler: inFleetVehicleCommandHandler}
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

func (ifvc *InstallVehicleController) Install(w http.ResponseWriter, r *http.Request) {
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

	var schema InstallVehicleUnmarshal
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

	command := install_vehicle_command.NewInstallVehicleCommand(schema.DeviceSerialNumber, schema.ChassisNumber)
	err = ifvc.installVehicleCommandHandler.Handle(command)

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
