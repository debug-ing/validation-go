package main

import (
	"errors"
	"fmt"

	"github.com/debug-ing/validation-go/pkg/validation"
)

type UpdateIAMUserRequest struct {
	ID    string `validate:"required,uuid"`
	Name  string `validate:"required,minlength=10"`
	LName string `validate:"required,minlength=10"`
	Age   int    `validate:"required,numeric,minmax=18-100"`
	Nike  string `validate:"required,between=1-10"`
	Test  string `validate:"contains=hello"`
}

func main() {
	req := UpdateIAMUserRequest{
		ID:    "d89cb260-36ff-4206-8979-6c1ce693ab28",
		Name:  "ddsffddddd",
		LName: "ddsffddddd",
		Age:   29,
		Nike:  "sdfkjdf",
		Test:  "hell1o",
	}
	validation.AddCustomValidator("test", "%s vard kon", validateRequired)
	err := validation.ValidateStruct(req)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}
	//
}

func validateRequired(value string, errorMsg string) error {
	if value == "" {
		fmt.Println(errorMsg)
		return errors.New(errorMsg)
	}
	return nil
}
