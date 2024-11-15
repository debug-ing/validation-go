package validation

import (
	"errors"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func validateRequired(value string, errorMsg string) error {
	if value == "" {
		return errors.New(errorMsg)
	}
	return nil
}

func validationIsUUID(value string, errorMsg string) error {
	_, err := uuid.Parse(value)
	if err != nil {
		return errors.New(errorMsg)
	}
	return nil
}

func validateMinLength(value string, min int, errorMsg string) error {
	if len(value) < min {
		return errors.New(errorMsg)
	}
	return nil
}
func validateIsEmail(value string, errorMsg string) error {
	if !strings.Contains(value, "@") {
		return errors.New(errorMsg)
	}
	return nil
}
func validateIsNumber(value string, errorMsg string) error {
	if _, err := strconv.Atoi(value); err != nil {
		return errors.New(errorMsg)
	}
	return nil
}

func validationBetween(value string, min, max int, errorMsg string) error {
	if len(value) < min || len(value) > max {
		return errors.New(errorMsg)
	}
	return nil
}

func validationMinMaxNumber(value string, min, max int, errorMsg string) error {
	num, err := strconv.Atoi(value)
	if err != nil {
		return errors.New(errorMsg)
	}

	if num < min || num > max {
		return errors.New(errorMsg)
	}
	return nil
}
