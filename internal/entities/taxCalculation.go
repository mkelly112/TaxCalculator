package entities

// Enhanced Brackets add info to a bracket based on salary under calculation to show the actual value within the bracket, and the tax of the bracket
type EnhancedBrackets struct {
	TaxBracket    TaxBracket `json:"tax_bracket"`
	BracketActual float64    `json:"actual"`
	BracketTax    float64    `json:"tax"`
}

// TaxCalculation represents the response for calculation of Tax Brackets including a total amount
type TaxCalculation struct {
	TotalTax         float64          `json:"total"`
	EnhancedBrackets EnhancedBrackets `json:"tax_brackets"`
}
