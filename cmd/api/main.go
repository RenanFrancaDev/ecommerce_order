package main

import (
	"ecommerce_order/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	http.RegisterRoutes(r)
	r.Run(":8080")
}