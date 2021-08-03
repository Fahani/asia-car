package get_telemetries_by_chassis_number_query

type GetTelemetriesByChassisNumberQuery struct {
	chassisNumber string
}

func NewGetTelemetriesByChassisNumberQuery(chassisNumber string) GetTelemetriesByChassisNumberQuery {
	return GetTelemetriesByChassisNumberQuery{chassisNumber: chassisNumber}
}

func (gtbcnq *GetTelemetriesByChassisNumberQuery) GetChassisNumber() string {
	return gtbcnq.chassisNumber
}
