package main

import (
	"log"
	"os"

	"github.com/sos/auth/be/go/init-go-gin/internal/app"
)

func main() {
	engine := app.New()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
