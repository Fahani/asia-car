package custom_error

import (
	"errors"
	"fmt"
)

type InstallVehicleError struct {
	errorMessage error
}

func (e *InstallVehicleError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func InstallVehicle(errorMessage string) *InstallVehicleError {
	e := new(InstallVehicleError)
	e.errorMessage = errors.New(errorMessage)
	return e
}
