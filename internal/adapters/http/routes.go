package http

import (
	"github.com/gin-gonic/gin"
)
func RegisterOrderRoutes(r *gin.Engine, h *OrderHandler) {
	r.POST("/orders", h.PlaceOrder)
}