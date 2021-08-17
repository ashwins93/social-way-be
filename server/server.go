package main

import (
	"net/http"
	"os"

	"github.com/ashwins93/social-way-be/config"
	"github.com/ashwins93/social-way-be/model"
	"github.com/ashwins93/social-way-be/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	db := config.GetConnection()

	db.AutoMigrate(model.Models...)

	routingService := routes.NewService(r, db)

	routingService.SetupV1Routes()

	r.Run(":" + port)
}
