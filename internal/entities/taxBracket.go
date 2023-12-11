package entities

// TaxBracket contains info on min, max, and rate for Canadian Federal Tax Brackets
type TaxBracket struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Rate float64 `json:"rate"`
}

// TaxBrackets represents an array of TaxBrackets, returned by the tax calculator
type TaxBrackets struct {
	TaxBrackets []TaxBracket `json:"tax_brackets"`
}
