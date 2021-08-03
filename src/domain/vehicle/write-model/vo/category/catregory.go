package category

import (
	"fmt"
	"github.com/fahani/asia-car/src/domain/common/custom-error"
	"regexp"
)

const validRegExp string = "^[MNEHCDIJSRFGPULWOX][BCDWVLSTFJXPQZEMRHYNGK][MNCABD][RNDQHIECLSABMFVZUX]$"

type Category struct {
	value string
}

func FromValue(value string) (Category, error) {
	err := assertValidValue(value)
	if err != nil {
		return Category{}, err
	}
	return Category{value: value}, nil
}

func assertValidValue(value string) error {
	re := regexp.MustCompile(validRegExp)
	match := re.MatchString(value)
	if !match {
		return custom_error.InvalidArgument(fmt.Sprintf("Invalid Category %s, the number should a valid ACRISS code", value))
	}
	return nil
}

func (n Category) GetValue() string {
	return n.value
}
