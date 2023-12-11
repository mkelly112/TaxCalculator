package helpers

import (
	"github.com/go-playground/validator/v10"
)

// GetTaxBracketParams are the parameters used to retrieve the tax bracket
type GetTaxBracketParams struct {
	Year string `form:"year" binding:"required,numeric,len=4"`
}

// GetCalculateTaxParams are the parameters used to calculate federal tax
type GetCalculateTaxParams struct {
	Year   string `form:"year" binding:"required,numeric,len=4"`
	Salary string `form:"salary" binding:"required,numeric"`
}

// IsValidYear checks if the input year is valid
func IsValidYear(year string) bool {
	validYears := []string{"2019", "2020", "2021", "2022"}
	for _, validYear := range validYears {
		if year == validYear {
			return true
		}
	}
	return false
}

// IsValidSalary checks if the salary is valid (non-negative)
func IsValidSalary(salary float64) bool {
	return salary < 0
}

// GetValidationError returns the validation error message based on the validation error.
func GetValidationError(ve validator.ValidationErrors) string {
	var yearErr, salErr string
	// go through cases to determine the validation error
	for _, e := range ve {
		switch e.Field() {
		case "Year":
			switch e.Tag() {
			case "required":
				yearErr = "Year is required"
			case "numeric":
				yearErr = "Year must be a 4 digit numeric value"
			case "len":
				yearErr = "Year must be exactly 4 digits"
			default:
				yearErr = "Inavalid Year"
			}
		case "Salary":
			switch e.Tag() {
			case "required":
				salErr = "Salary is required"
			case "numeric":
				salErr = "Salary must be a numeric value"
			default:
				salErr = "Invalid Salary"
			}
		default:
			yearErr = "Invalid Year or Salary"
		}
	}

	// Combine errors to display any issues in the same message
	errMsg := yearErr + "/n" + salErr

	return errMsg
}
