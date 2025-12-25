package main

import (
	"log"
	"path/filepath"
	"phongtran/go-social/golang-social/internal/app"
	"phongtran/go-social/golang-social/internal/config"
	"phongtran/go-social/golang-social/internal/utils"

	"github.com/joho/godotenv"
)

func main() {
	rootDir := utils.GetWorkingDir()

	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		log.Fatal("⛔ Error loading .env file")
	}

	log.Println("✅ Loading .env file")

	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize configuration
	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
