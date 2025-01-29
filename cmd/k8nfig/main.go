package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/podummy/config"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app_name":    cfg.App.Name,
			"app_author":  cfg.App.Author,
			"app_version": cfg.App.Version,
		})
	})

	r.Run(":8080")
}
