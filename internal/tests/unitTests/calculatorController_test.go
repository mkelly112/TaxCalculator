package tests

import (
	"TaxCalculator/internal/controllers"
	"TaxCalculator/internal/controllers/helpers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCalculateTax(t *testing.T) {
	// Set up mocks required for test.
	router := gin.New()
	taxBracketService := &mockTaxBracketService{}
	taxCalculationService := &mockTaxCalculationService{}
	calculatorController := controllers.NewCalculatorController(taxBracketService, taxCalculationService)
	router.Handle(http.MethodGet, "/tax-calculator", calculatorController.CalculateTax)

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tax-calculator?year=2022&salary=70000", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
		}

		var response helpers.CalculatorResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		expectedTotal := 11589.17
		if response.TaxCalculation.TotalTax != expectedTotal {
			t.Errorf("Expected total tax amount %f, but got %f", expectedTotal, response.TaxCalculation.TotalTax)
		}
	})

	t.Run("Invalid URL", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/pokemon", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, rec.Code)
		}
	})

	t.Run("Invalid Tax Year", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tax-calculator?year=2023&salary=70000", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Invalid Salary", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tax-calculator?year=2022&salary=turtle", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Negative Salary", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tax-calculator?year=2022&salary=-70000", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
		}
	})

}
