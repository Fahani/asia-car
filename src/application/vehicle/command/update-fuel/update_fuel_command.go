package update_fuel_command

type UpdateFuelCommand struct {
	deviceSerialNumber string
	fuel               int
	updateType         string
}

func NewUpdateFuelCommand(fuel int, deviceSerialNumber, updateType string) UpdateFuelCommand {
	return UpdateFuelCommand{
		deviceSerialNumber: deviceSerialNumber,
		fuel:               fuel,
		updateType:         updateType,
	}
}

func (ufc *UpdateFuelCommand) GetDeviceSerialNumber() string {
	return ufc.deviceSerialNumber
}

func (ufc *UpdateFuelCommand) GetFuel() int {
	return ufc.fuel
}

func (ufc *UpdateFuelCommand) GetUpdateType() string {
	return ufc.updateType
}
