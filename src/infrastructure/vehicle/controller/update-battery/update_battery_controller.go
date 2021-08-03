package update_battery_controller

import (
	"encoding/json"
	"fmt"
	"github.com/fahani/asia-car/src/application/vehicle/command/update-battery"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
	"net/http"
)

type UpdateBatteryController struct {
	updateBatteryCommandHandler *update_battery_command.UpdateBatteryLevelCommandHandler
}
type UpdateBatteryUnmarshal struct {
	DeviceSerialNumber string `json:"device_serial_number"`
	Battery            int    `json:"battery"`
}

const jsonSchemaValidation = `{
  "type":"object",
  "required":[
    "device_serial_number",
	"battery"
  ],
  "properties":{
    "device_serial_number":{
      "type":"string"
    },
	"battery":{
      "type":"integer"
    }
  }
}`

func NewUpdateBatteryController(updateBatteryCommandHandler *update_battery_command.UpdateBatteryLevelCommandHandler) UpdateBatteryController {
	return UpdateBatteryController{updateBatteryCommandHandler: updateBatteryCommandHandler}
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

func (ifvc *UpdateBatteryController) UpdateBattery(w http.ResponseWriter, r *http.Request) {
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

	var schema UpdateBatteryUnmarshal
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

	command := update_battery_command.NewUpdateBatteryCommand(schema.DeviceSerialNumber, schema.Battery)
	err = ifvc.updateBatteryCommandHandler.Handle(command)

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
