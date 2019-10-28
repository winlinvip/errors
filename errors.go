package errors

import "errors"

func Errorf(format string, args... interface{}) error {
	return errors.New("mod replaced")
}
