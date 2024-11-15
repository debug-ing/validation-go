package validation

import (
	"testing"
)

func TestValidateRequired(t *testing.T) {
	tests := []struct {
		value    string
		errorMsg string
		wantErr  bool
	}{
		{"", "This field is required", true},       // Should return error
		{"Hello", "This field is required", false}, // Should not return error
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			err := validateRequired(tt.value, tt.errorMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateRequired() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidationIsUUID(t *testing.T) {
	tests := []struct {
		value    string
		errorMsg string
		wantErr  bool
	}{
		{"invalid-uuid", "Invalid UUID", true},                          // Should return error
		{"f47ac10b-58cc-4372-a567-0e02b2c3d479", "Invalid UUID", false}, // Should not return error
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			err := validationIsUUID(tt.value, tt.errorMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationIsUUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMinLength(t *testing.T) {
	tests := []struct {
		value    string
		min      int
		errorMsg string
		wantErr  bool
	}{
		{"short", 10, "Must be at least 10 characters", true},      // Should return error
		{"longenough", 5, "Must be at least 10 characters", false}, // Should not return error
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			err := validateMinLength(tt.value, tt.min, tt.errorMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateMinLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateIsEmail(t *testing.T) {
	tests := []struct {
		value    string
		errorMsg string
		wantErr  bool
	}{
		{"invalid-email", "Invalid email", true},     // Should return error
		{"test@example.com", "Invalid email", false}, // Should not return error
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			err := validateIsEmail(tt.value, tt.errorMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateIsEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateIsNumber(t *testing.T) {
	tests := []struct {
		value    string
		errorMsg string
		wantErr  bool
	}{
		{"not-a-number", "Must be a number", true}, // Should return error
		{"12345", "Must be a number", false},       // Should not return error
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			err := validateIsNumber(tt.value, tt.errorMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateIsNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// func TestValidationBetween(t *testing.T) {
// 	tests := []struct {
// 		value    string
// 		min      int
// 		max      int
// 		errorMsg string
// 		wantErr  bool
// 	}{
// 		{"short", 5, 10, "Length must be between 5 and 10", true},    // Should return error
// 		{"perfect", 5, 10, "Length must be between 5 and 10", false}, // Should not return error
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.value, func(t *testing.T) {
// 			err := validationBetween(tt.value, tt.min, tt.max, tt.errorMsg)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("validationBetween() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func TestValidationMinMaxNumber(t *testing.T) {
	tests := []struct {
		value    string
		min      int
		max      int
		errorMsg string
		wantErr  bool
	}{
		{"50", 10, 100, "Number must be between 10 and 100", false}, // Should not return error
		{"200", 10, 100, "Number must be between 10 and 100", true}, // Should return error
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			err := validationMinMaxNumber(tt.value, tt.min, tt.max, tt.errorMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMinMaxNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
