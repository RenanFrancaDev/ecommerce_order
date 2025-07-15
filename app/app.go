package app

import (
	"ecommerce_order/internal/adapters/http"
	"ecommerce_order/internal/config"

	"github.com/gin-gonic/gin"
)
type App struct {
	cfg      *config.Config
	handlers *http.Handlers
	router   *gin.Engine
}
func NewApp() *App {
	return &App{}
}
func (a *App) BuildConfig() *App {
	a.cfg = config.Load()
	return a
}
func (a *App) BuildHandlers() *App {
	a.handlers = http.BuildHandlers(a.cfg)
	return a
}
func (a *App) BuildRouter() *App {
	a.router = gin.Default()
	return a
}
func (a *App) MapWebRoutes() *App {
	http.RegisterOrderRoutes(a.router, a.handlers.Order)
	return a
}
func (a *App) Run() {
	a.router.Run(":8080")
}