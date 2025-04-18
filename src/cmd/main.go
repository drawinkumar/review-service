package main

import (
	"log"

	"example.com/review/v2/config"
	"example.com/review/v2/controller"
)

func main() {
	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// run the app
	app := controller.New(cfg)
	if err := app.Run(); err != nil {
		log.Fatalf("App error: %v", err)
	}
}
