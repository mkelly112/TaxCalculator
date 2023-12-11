package services

import (
	"TaxCalculator/internal/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ITaxBracketService is the interface for retrieving tax brackets
type ITaxBracketService interface {
	GetTaxBracket(year string) (*entities.TaxBrackets, error)
}

// taxBracketService will require url endpoint for retrieving tax brackets
type taxBracketService struct {
	bracketCalculatorURL string
}

// NewTaxBracketService creates a new instance of the taxBracketService
func NewTaxBracketService(bracketCalculatorURL string) ITaxBracketService {
	return &taxBracketService{
		bracketCalculatorURL: bracketCalculatorURL,
	}
}

// GetTaxBracket retrieves the tax brackets for the supplied year from the tax bracket api
func (s *taxBracketService) GetTaxBracket(year string) (*entities.TaxBrackets, error) {
	// maxRetries is the maximum number of retries to aquire tax bracket info from the api
	maxRetries := 3
	// retryInterval is the duration between retries in seconds
	retryInterval := 2 * time.Second

	for retry := 0; retry < maxRetries; retry++ {
		resp, err := http.Get(fmt.Sprintf(s.bracketCalculatorURL + year))
		if err != nil {
			time.Sleep(retryInterval)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			time.Sleep(retryInterval)
			continue
		}

		var brackets entities.TaxBrackets
		err = json.NewDecoder(resp.Body).Decode(&brackets)
		if err != nil {
			time.Sleep(retryInterval)
			continue
		}

		return &brackets, nil
	}

	return nil, fmt.Errorf("Failed to retrieve tax bracket information, please try again")
}
