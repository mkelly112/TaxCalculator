package controllers

import (
	"TaxCalculator/internal/controllers/helpers"
	"TaxCalculator/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BracketController struct {
	bracketService services.ITaxBracketService
}

// NewBracketController creates a new instance of BracketController with the given bracket service
func NewBracketController(bracketService services.ITaxBracketService) *BracketController {
	return &BracketController{
		bracketService: bracketService,
	}
}

// GetTaxBracket returns the tax brackets for a specified tax year
func (c *BracketController) GetTaxBracket(ctx *gin.Context) {
	var qp helpers.GetTaxBracketParams
	if err := ctx.ShouldBindQuery(&qp); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMsg := helpers.GetValidationError(ve)
			helpers.BadRequest(ctx, errorMsg)
			return
		}

		helpers.BadRequest(ctx, "Invalid query parameters")
		return
	}

	// Get year from query parameters
	year := ctx.Query("year")

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

	response := helpers.BracketsResponse{
		TaxBrackets: *brackets,
	}
	helpers.OK(ctx, response)
}
