package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nepile/api-monitoring/config"
	"github.com/nepile/api-monitoring/database"
	"github.com/nepile/api-monitoring/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	database.Connect(cfg.DatabaseURL)

	r := gin.Default()
	routes.Setup(r, cfg)

	log.Println("listening on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
