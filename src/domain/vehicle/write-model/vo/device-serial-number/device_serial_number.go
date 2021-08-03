package device_serial_number

type DeviceSerialNumber struct {
	value string
}

func FromValue(value string) DeviceSerialNumber {
	return DeviceSerialNumber{value: value}
}

func (n DeviceSerialNumber) GetValue() string {
	return n.value
}
