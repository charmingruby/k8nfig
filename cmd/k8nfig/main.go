package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/podummy/config"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	cfg, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: Failed to load configuration: %s", err.Error()))
	}

	slog.Info(fmt.Sprintf("App Name: %s", cfg.App.Name))
	slog.Info(fmt.Sprintf("App Author: %s", cfg.App.Author))
	slog.Info(fmt.Sprintf("App Version: %s", cfg.App.Version))
}
