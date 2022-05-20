package main

import (
	"log"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
	}

	// Run
	app.Run(cfg)
}
