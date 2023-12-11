package main

import (
	"TaxCalculator/internal/controllers"
	"TaxCalculator/internal/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// initApp initializes services
// Read environment variables, such as the bracket calculator url
func initApp() (*gin.Engine, error) {
	// Return error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Get bracketURL from .env file
	bracketCalculatorURL := os.Getenv("BRACKET_CALCULATOR_URL")

	// Set a default url if not found in .env
	if bracketCalculatorURL == "" {
		bracketCalculatorURL = "http://localhost:5000/tax-calculator/tax-year/"
	}

	// initalize tax bracket service and controller
	taxBracketService := services.NewTaxBracketService(bracketCalculatorURL)
	taxBracketController := controllers.NewBracketController(taxBracketService)

	// Create default Gin router
	router := gin.Default()

	// Register routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Tax Calculator Home"})
	})

	router.GET("/brackets", func(c *gin.Context) {
		taxBracketController.GetTaxBracket(c)
	})

	// Error handler for undefined routes
	router.NoRoute(func(c *gin.Context) {
		// Return 404 when the route is undefined
		c.JSON(404, gin.H{"message": "Route not found"})
	})

	return router, nil
}
