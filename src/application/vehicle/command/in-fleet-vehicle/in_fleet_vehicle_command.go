package in_fleet_vehicle_command

type InFleetVehicleCommand struct {
	inFleetDate   string
	chassisNumber string
}

func NewInFleetVehicleCommand(inFleetDate, chassisNumber string) InFleetVehicleCommand {
	return InFleetVehicleCommand{inFleetDate: inFleetDate, chassisNumber: chassisNumber}
}

func (ifv *InFleetVehicleCommand) GetInFleetDate() string {
	return ifv.inFleetDate
}

func (ifv *InFleetVehicleCommand) GetChassisNumber() string {
	return ifv.chassisNumber
}
