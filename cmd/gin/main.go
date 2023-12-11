package main

import (
	"log"
)

func main() {
	// initialize services, controllers, etc, obtaining the final router for the app
	router, err := initApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// port for the application
	port := ":8080"

	// start the app on port above
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
