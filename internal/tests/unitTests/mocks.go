package tests

import "TaxCalculator/internal/entities"

// Define a mock tax calculation service
type mockTaxCalculationService struct{}

// Create some mocks for tax calculation
func (m *mockTaxCalculationService) CalculateTax(taxBrackets *entities.TaxBrackets, salary float64) (*entities.TaxCalculation, error) {
	var taxCalculation entities.TaxCalculation
	if salary == 1 {
		taxCalculation.TotalTax = 0.15
		taxCalculation.EffectiveRate = 0.15
	} else if salary == 70000 {
		taxCalculation.TotalTax = 11589.17
		taxCalculation.EffectiveRate = 0.1656
	} else {
		taxCalculation.TotalTax = 2.55
		taxCalculation.EffectiveRate = 0.15
	}
	return &taxCalculation, nil
}

// Define a mock tax bracket service
type mockTaxBracketService struct{}

func (m *mockTaxBracketService) GetTaxBracket(year string) (*entities.TaxBrackets, error) {
	// mock 2022 tax brackets
	taxBrackets := &entities.TaxBrackets{
		TaxBrackets: []entities.TaxBracket{
			{
				Min:  0,
				Max:  50197,
				Rate: 0.15,
			},
			{
				Min:  50197,
				Max:  100392,
				Rate: 0.205,
			},
			{
				Min:  100392,
				Max:  155625,
				Rate: 0.26,
			},
			{
				Min:  155625,
				Max:  221708,
				Rate: 0.29,
			},
			{
				Min:  221708,
				Max:  0,
				Rate: 0.33,
			},
		},
	}

	return taxBrackets, nil
}
