package chassis_nbr

import (
	"fmt"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"regexp"
)

const validRegExp string = "^[a-zA-Z0-9]{17}$"

type ChassisNbr struct {
	value string
}

func FromValue(value string) (ChassisNbr, error) {
	err := assertValidValue(value)
	if err != nil {
		return ChassisNbr{}, err
	}
	return ChassisNbr{value: value}, err
}

func assertValidValue(value string) error {
	re := regexp.MustCompile(validRegExp)
	match := re.MatchString(value)
	if !match {
		return custom_error.InvalidArgument(fmt.Sprintf("Invalid Chassis Number %s, the number should be 17 digits alphanumeric", value))
	}
	return nil
}

func (n ChassisNbr) GetValue() string {
	return n.value
}
