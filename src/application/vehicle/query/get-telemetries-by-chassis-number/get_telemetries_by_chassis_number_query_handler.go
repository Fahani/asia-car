package get_telemetries_by_chassis_number_query

import (
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/repository/read"
	"github.com/fahani/asia-car/src/domain/vehicle/write-model/vo/chassis-nbr"
)

type GetTelemetriesByChassisNumberQueryHandler struct {
	vehicleReadRepository vehicle_read_repostiroy.VehicleReadRepository
}

func NewGetTelemetriesByChassisNumberQueryHandler(vehicleReadRepository vehicle_read_repostiroy.VehicleReadRepository) *GetTelemetriesByChassisNumberQueryHandler {
	return &GetTelemetriesByChassisNumberQueryHandler{vehicleReadRepository: vehicleReadRepository}
}

func (gtbcnqh *GetTelemetriesByChassisNumberQueryHandler) Handle(query GetTelemetriesByChassisNumberQuery) (GetTelemetriesByChassisNumberQueryResponse, error) {
	chassisNumberVO, err := chassis_nbr.FromValue(query.GetChassisNumber())

	if err != nil {
		return GetTelemetriesByChassisNumberQueryResponse{}, err
	}

	vehicle, err := gtbcnqh.vehicleReadRepository.GetVehicleByChassisNbr(chassisNumberVO)

	if err != nil {
		return GetTelemetriesByChassisNumberQueryResponse{}, err
	}

	getTelemetriesBySerialNumberQueryResponse := NewGetTelemetriesBySerialNumberQueryResponse(vehicle)

	return getTelemetriesBySerialNumberQueryResponse, nil
}
