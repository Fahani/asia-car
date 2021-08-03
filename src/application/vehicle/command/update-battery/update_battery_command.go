package update_battery_command

type UpdateBatteryCommand struct {
	deviceSerialNumber string
	batteryLevel       int
}

func NewUpdateBatteryCommand(deviceSerialNumber string, batteryLevel int) UpdateBatteryCommand {
	return UpdateBatteryCommand{
		deviceSerialNumber: deviceSerialNumber,
		batteryLevel:       batteryLevel,
	}
}

func (ubc *UpdateBatteryCommand) GetDeviceSerialNumber() string {
	return ubc.deviceSerialNumber
}
func (ubc *UpdateBatteryCommand) GetBatteryLevel() int {
	return ubc.batteryLevel
}
