package entities

// Enhanced Brackets add info to a bracket based on salary under calculation to show the actual value within the bracket, and the tax of the bracket
type EnhancedBracket struct {
	Min           float64 `json:"min"`
	Max           float64 `json:"max"`
	Rate          float64 `json:"rate"`
	BracketActual float64 `json:"actual"`
	BracketTax    float64 `json:"tax"`
}

// TaxCalculation represents the response for calculation of Tax Brackets including a total amount
type TaxCalculation struct {
	TotalTax         float64           `json:"total"`
	EffectiveRate    float64           `json:"effective_rate"`
	EnhancedBrackets []EnhancedBracket `json:"tax_brackets"`
}
