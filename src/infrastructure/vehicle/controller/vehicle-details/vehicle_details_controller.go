package vehicle_details_controller

import (
	"encoding/json"
	"fmt"
	"github.com/fahani/asia-car/src/application/vehicle/query/get-telemetries-by-chassis-number"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
	"net/http"
)

type VehicleDetailsController struct {
	getTelemetriesCommandHandler *get_telemetries_by_chassis_number_query.GetTelemetriesByChassisNumberQueryHandler
}

type VehicleDetailsUnmarshal struct {
	ChassisNumber string `json:"chassis_number"`
}

const jsonSchemaValidation = `{
  "type":"object",
  "required":[
    "chassis_number"
  ],
  "properties":{
    "chassis_number":{
      "type":"string"
    }
  }
}`

func NewVehicleDetailsController(getTelemetriesCommandHandler *get_telemetries_by_chassis_number_query.GetTelemetriesByChassisNumberQueryHandler) VehicleDetailsController {
	return VehicleDetailsController{getTelemetriesCommandHandler: getTelemetriesCommandHandler}
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

func (ifvc *VehicleDetailsController) VehicleDetails(w http.ResponseWriter, r *http.Request) {
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

	var schema VehicleDetailsUnmarshal
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

	query := get_telemetries_by_chassis_number_query.NewGetTelemetriesByChassisNumberQuery(schema.ChassisNumber)
	response, err := ifvc.getTelemetriesCommandHandler.Handle(query)

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

	data, _ := json.Marshal(response)
	w.WriteHeader(200)
	_, err = w.Write(data)
	if nil != err {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	return
}
