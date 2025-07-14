package main

import (
	"ecommerce_order/internal/adapters/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	http.RegisterRoutes(r)
	r.Run(":8080")
}