package main

import (
	"fmt"

	"github.com/debug-ing/validation-go/pkg/validation"
)

type UpdateIAMUserRequest struct {
	ID    string `validate:"required,uuid"`
	Name  string `validate:"required,minlength=10"`
	LName string `validate:"required,minlength=10"`
	Age   int    `validate:"required,numeric,minmax=18-100"`
	Nike  string `validate:"required,between=1-10"`
}

func main() {
	req := UpdateIAMUserRequest{
		ID:    "d89cb260-36ff-4206-8979-6c1ce693ab28",
		Name:  "ddsffddddd",
		LName: "ddsffddddd",
		Age:   29,
		Nike:  "sdfkjdf",
	}
	//
	err := validation.ValidateStruct(req)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}
}
