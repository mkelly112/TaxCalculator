package controllers

import (
	"TaxCalculator/internal/controllers/helpers"
	"TaxCalculator/internal/services"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CalculatorController struct {
	bracketService       services.ITaxBracketService
	taxCalculatorService services.ITaxCalculatorService
}

// NewCalculatorController creates a new instance of CalculatorController with the given bracket and calculator services
func NewCalculatorController(bracketService services.ITaxBracketService, taxCalculatorService services.ITaxCalculatorService) *CalculatorController {
	return &CalculatorController{
		bracketService:       bracketService,
		taxCalculatorService: taxCalculatorService,
	}
}

// CalculateTax calculates the taxes for each bracket, as well as total tax and effective rate for a specific tax year and salary
func (c *CalculatorController) CalculateTax(ctx *gin.Context) {
	// Validate query parameters
	var qp helpers.GetCalculateTaxParams
	if err := ctx.ShouldBindQuery(&qp); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMsg := helpers.GetValidationError(ve)
			helpers.BadRequest(ctx, errorMsg)
			return
		}
	}

	// Get year,salary from query params
	year := ctx.Query("year")
	salaryStr := ctx.Query("salary")

	// Validate/convert salary input
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		helpers.InternalServerError(ctx, "Invalid salary")
		return
	}

	// Validate the valid tax year input
	if !helpers.IsValidYear(year) {
		helpers.BadRequest(ctx, "Invalid tax year. Please select one of 2019,2020,2021,2022.")
		return
	}

	// Get the tax brackets from the bracketService for the given year
	brackets, err := c.bracketService.GetTaxBracket(year)
	if err != nil {
		helpers.InternalServerError(ctx, "Failed to obtain tax brackets")
		return
	}

	// Get the tax calculations from the calculator service
	taxCalculation, err := c.taxCalculatorService.CalculateTax(brackets, salary)
	if err != nil {
		helpers.InternalServerError(ctx, "Failed to calculate tax")
		return
	}

	response := helpers.CalculatorResponse{
		TaxCalculation: *taxCalculation,
	}
	helpers.OK(ctx, response)
}
