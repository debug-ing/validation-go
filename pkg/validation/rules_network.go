package validation

import (
	"errors"
	"regexp"
)

func validationIsURL(value string, errorMsg string) error {
	urlRegex := `^(http|https)://[a-zA-Z0-9\-.]+\.[a-zA-Z]{2,}(/\S*)?$`
	if matched, _ := regexp.MatchString(urlRegex, value); !matched {
		return errors.New(errorMsg)
	}
	return nil
}

func validationIsIPv4(value string, errorMsg string) error {
	ipRegex := `^(\d{1,3}\.){3}\d{1,3}$`
	if matched, _ := regexp.MatchString(ipRegex, value); !matched {
		return errors.New(errorMsg)
	}
	return nil
}

func validationIsIPv6(value string, errorMsg string) error {
	ipRegex := `^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$`
	if matched, _ := regexp.MatchString(ipRegex, value); !matched {
		return errors.New(errorMsg)
	}
	return nil
}

func validationIsIP(value string, errorMsg string) error {
	if err := validationIsIPv4(value, errorMsg); err != nil {
		if err := validationIsIPv6(value, errorMsg); err != nil {
			return errors.New(errorMsg)
		}
	}
	return nil
}
