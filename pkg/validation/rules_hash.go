package validation

import "errors"

func validationMd5(value string, errorMsg string) error {
	if len(value) != 32 {
		return errors.New(errorMsg)
	}
	return nil
}
