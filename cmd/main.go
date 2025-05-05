package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/usama1031/go-sms-verify/api"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	app := api.Config{Router: router}

	app.Routes()

	router.Run(":" + port)
}
