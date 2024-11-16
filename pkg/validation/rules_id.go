package validation

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

func validationIsMongoID(value string, errorMsg string) error {
	mongoIDRegex := `^[0-9a-fA-F]{24}$`
	if matched, _ := regexp.MatchString(mongoIDRegex, value); !matched {
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

func validationIsULID(value string, errorMsg string) error {
	_, err := ulid.Parse(value)
	if err != nil {
		return errors.New(errorMsg)
	}
	return nil
}
