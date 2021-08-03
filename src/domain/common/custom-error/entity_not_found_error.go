package custom_error

import (
	"errors"
	"fmt"
)

type EntityNotFoundError struct {
	errorMessage error
}

func (e *EntityNotFoundError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func EntityNotFound(errorMessage string) *EntityNotFoundError {
	e := new(EntityNotFoundError)
	e.errorMessage = errors.New(errorMessage)
	return e
}
