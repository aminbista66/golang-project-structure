package main

import (
	"log"
	"myapp/internal/config"
	"myapp/internal/app"
)

func main() {
	cfg := config.Load()
	container, err := app.Initialize(cfg)
	if err != nil {
		log.Fatalf("failed to initialize container: %v", err)
	}
	log.Printf("🚀 Server running on port %s", cfg.Server.Port)
	if err := container.Router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
