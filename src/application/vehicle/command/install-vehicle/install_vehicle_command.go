package install_vehicle_command

type InstallVehicleCommand struct {
	deviceSerialNumber string
	chassisNumber      string
}

func NewInstallVehicleCommand(deviceSerialNumber, chassisNumber string) InstallVehicleCommand {
	return InstallVehicleCommand{
		deviceSerialNumber: deviceSerialNumber,
		chassisNumber:      chassisNumber,
	}
}

func (ivc *InstallVehicleCommand) GetDeviceSerialNumber() string {
	return ivc.deviceSerialNumber
}
func (ivc *InstallVehicleCommand) GetChassisNumber() string {
	return ivc.chassisNumber
}
