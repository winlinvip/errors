package errors

import (
	"errors"
	"fmt"
)

func Errorf(format string, args... interface{}) error {
	return errors.New(fmt.Sprintf("[mod replace core] %v", fmt.Sprintf(format, args...)))
}
