package helpers

import (
	"TaxCalculator/internal/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BracketsResponse represents the response for GetBrackets endpoint
type BracketsResponse struct {
	TaxBrackets entities.TaxBrackets `json:"tax_brackets"`
}

// CalculatorResponse represents the response for the CalculateTax endpoint
type CalculatorResponse struct {
	TaxCalculation entities.TaxCalculation `json:"tax_calculation"`
}

// APIError represents the JSON response for API errors
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// BadRequest sends a Bad Request response with the provided error message.
func BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, APIError{
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

// InternalServerError sends an Internal Server Error response with the provided error message.
func InternalServerError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, APIError{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

// OK sends a success response with the provided data.
func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
