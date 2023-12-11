package services

import (
	"TaxCalculator/internal/entities"
	"math"
)

// ITaxCalculator Service is the interface for calculating taxes owed
type ITaxCalculatorService interface {
	CalculateTax(taxBrackets *entities.TaxBrackets, salary float64) (*entities.TaxCalculation, error)
}

type taxCalculatorService struct{}

// NewTaxCalculator creates a new instance of the taxCalculatorService
func NewTaxCalculator() ITaxCalculatorService {
	return &taxCalculatorService{}
}

// CalculateTax retrieves the tax brackets for a supplied year and calculates the tax owed for each bracket
// as well as calculating the total tax and effective rate
func (s *taxCalculatorService) CalculateTax(taxBrackets *entities.TaxBrackets, salary float64) (*entities.TaxCalculation, error) {
	var enhancedBrackets []entities.EnhancedBracket
	var totalTax float64
	totalTax = 0

	for _, bracket := range taxBrackets.TaxBrackets {
		var enhancedBracket entities.EnhancedBracket
		// copy bracket info to the enhanced bracket
		enhancedBracket.Min = bracket.Min
		enhancedBracket.Max = bracket.Max
		enhancedBracket.Rate = bracket.Rate
		// calculate the bracket taxable amount
		if bracket.Max != 0 && salary > bracket.Max {
			enhancedBracket.BracketActual = bracket.Max - bracket.Min
		} else if salary > bracket.Min {
			enhancedBracket.BracketActual = salary - bracket.Min
		} else {
			enhancedBracket.BracketActual = 0
		}
		// calculate bracket tax
		enhancedBracket.BracketTax = math.Round(100*(enhancedBracket.BracketActual*enhancedBracket.Rate)) / 100
		totalTax += enhancedBracket.BracketTax
		enhancedBrackets = append(enhancedBrackets, enhancedBracket)
	}

	result := &entities.TaxCalculation{
		TotalTax:         totalTax,
		EffectiveRate:    totalTax / salary,
		EnhancedBrackets: enhancedBrackets,
	}

	return result, nil
}
