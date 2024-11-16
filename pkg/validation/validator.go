package validation

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var validators = map[string]interface{}{
	"required":    validateRequired,
	"minlength":   validateMinLength,
	"uuid":        validationIsUUID,
	"ulid":        validationIsULID,
	"bic":         validationIsBIC,
	"ethaddress":  validationIsEthereumAddress,
	"btcaddress":  validationIsBtcAddress,
	"mongoID":     validationIsMongoID,
	"email":       validateIsEmail,
	"numeric":     validateIsNumber,
	"boolean":     validateIsBoolean,
	"contains":    validationIsContains,
	"notcontains": validationIsNotContains,
	"between":     validationBetween,
	"minmax":      validationMinMaxNumber,
}

var data = map[string]map[string]string{
	"required": {
		"error_msg": "%s field is required.",
	},
	"minlength": {
		"error_msg": "%s must be at least %d  characters long.",
	},
	"uuid": {
		"error_msg": "%s must be a valid UUID.",
	},
	"email": {
		"error_msg": "%s must be a valid email address.",
	},
	"numeric": {
		"error_msg": "%s must be a number.",
	},
	"between": {
		"error_msg": "%s must be between %d and %d characters long.",
	},
	"minmax": {
		"error_msg": "%s must be between %d and %d.",
	},
	"contains": {
		"error_msg": "%s must contain.",
	},
	"notcontains": {
		"error_msg": "%s must not contain.",
	},
}

func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fieldName := field.Name
		validateTag := field.Tag.Get("validate")
		data := strings.Split(validateTag, ",")
		for _, d := range data {
			parts := strings.SplitN(d, "=", 2)
			tagName := parts[0]
			var tagParam string
			if len(parts) > 1 {
				tagParam = parts[1]
			}
			if validator, exists := validators[tagName]; exists {
				switch validator := validator.(type) {
				case func(string, string) error:
					errorMsg := getMessage(tagName, fieldName)
					if err := validator(fmt.Sprint(value), errorMsg); err != nil {
						return err
					}
				case func(string, int, string) error:
					if param, err := strconv.Atoi(tagParam); err == nil {
						errorMsg := getMessage(tagName, fieldName, param)
						if err := validator(value.String(), param, errorMsg); err != nil {
							return err
						}
					} else {
						return fmt.Errorf("invalid parameter for tag %s", tagName)
					}
				case func(string, string, string) error:
					errorMsg := getMessage(tagName, fieldName)
					if err := validator(fmt.Sprint(value), tagParam, errorMsg); err != nil {
						return err
					}
				case func(string, int, int, string) error:
					params := strings.Split(tagParam, "-")
					if len(params) != 2 {
						return fmt.Errorf("invalid parameter for tag %s", tagName)
					}
					min, err := strconv.Atoi(params[0])
					if err != nil {
						return fmt.Errorf("invalid parameter for tag %s", tagName)
					}

					max, err := strconv.Atoi(params[1])
					if err != nil {
						return fmt.Errorf("invalid parameter for tag %s", tagName)
					}
					errorMsg := getMessage(tagName, fieldName, min, max)
					if err := validator(fmt.Sprint(value), min, max, errorMsg); err != nil {
						return err
					} else {

					}

				}
			}
		}
	}
	return nil
}

func AddCustomValidator(tagName, message string, fn interface{}) {
	validators[tagName] = fn
	data[tagName] = map[string]string{
		"error_msg": message,
	}
	fmt.Println(data)
}

func getMessage(tagName, fieldName string, args ...interface{}) string {
	errorMsgTemplate := data[tagName]["error_msg"]
	return fmt.Sprintf(errorMsgTemplate, append([]interface{}{fieldName}, args...)...)
}
