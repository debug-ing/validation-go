package validation

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
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

func validationIsULID(value string, errorMsg string) error {
	_, err := ulid.Parse(value)
	if err != nil {
		return errors.New(errorMsg)
	}
	return nil
}

func validationIsBIC(value string, errorMsg string) error {
	bicRegex := `^[A-Za-z]{4}[A-Za-z]{2}[A-Za-z0-9]{2}([A-Za-z0-9]{3})?$`
	if matched, _ := regexp.MatchString(bicRegex, value); !matched {
		return errors.New(errorMsg)
	}
	return nil
}

func validationIsEthereumAddress(value string, errorMsg string) error {
	ethRegex := `^0x[a-fA-F0-9]{40}$`
	if matched, _ := regexp.MatchString(ethRegex, value); !matched {
		return errors.New(errorMsg)
	}
	return nil
}

func validationIsBtcAddress(value string, errorMsg string) error {
	btcRegex := `^(1|3|bc1)[a-zA-HJ-NP-Z0-9]{25,39}$`
	if matched, _ := regexp.MatchString(btcRegex, value); !matched {
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
func validateIsBoolean(value string, errorMsg string) error {
	if _, err := strconv.ParseBool(value); err != nil {
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

func validationIsContains(value string, contains string, errorMsg string) error {
	if !strings.Contains(value, contains) {
		return errors.New(errorMsg)
	}
	return nil
}

func validationIsNotContains(value string, contains string, errorMsg string) error {
	if strings.Contains(value, contains) {
		return errors.New(errorMsg)
	}
	return nil
}

func validationIsMongoID(value string, errorMsg string) error {
	mongoIDRegex := `^[0-9a-fA-F]{24}$`
	if matched, _ := regexp.MatchString(mongoIDRegex, value); !matched {
		return errors.New(errorMsg)
	}
	return nil
}
