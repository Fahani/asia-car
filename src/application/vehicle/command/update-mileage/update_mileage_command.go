package update_mileage_command

type UpdateMileageCommand struct {
	deviceSerialNumber string
	mileage            int
}

func NewUpdateMileageCommand(deviceSerialNumber string, mileage int) UpdateMileageCommand {
	return UpdateMileageCommand{
		deviceSerialNumber: deviceSerialNumber,
		mileage:            mileage,
	}
}

func (ubc *UpdateMileageCommand) GetDeviceSerialNumber() string {
	return ubc.deviceSerialNumber
}
func (ubc *UpdateMileageCommand) GetMileageLevel() int {
	return ubc.mileage
}
