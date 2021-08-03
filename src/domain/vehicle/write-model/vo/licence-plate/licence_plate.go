package licence_plate

import (
	"fmt"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"regexp"
)

type LicencePlate struct {
	value string
}

const validRegExp string = "^[a-zA-Z0-9]+$"

func FromValue(value string) (LicencePlate, error) {
	err := assertValidValue(value)
	if err != nil {
		return LicencePlate{}, err
	}
	return LicencePlate{value: value}, err
}

func assertValidValue(value string) error {
	re := regexp.MustCompile(validRegExp)
	match := re.MatchString(value)
	if !match {
		return custom_error.InvalidArgument(fmt.Sprintf("Invalid Licence Plate %s, the number should be at least one digit alphanumeric", value))
	}
	return nil
}

func (n LicencePlate) GetValue() string {
	return n.value
}
