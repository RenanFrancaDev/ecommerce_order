package app

import (
	"ecommerce_order/internal/infrastructure/adapters/http"
	"ecommerce_order/internal/infrastructure/config"
	"ecommerce_order/internal/infrastructure/container"

	"github.com/gin-gonic/gin"
)

type App struct {
	cfg       *config.Config
	container *container.Container
	handlers  *http.Handlers
	router    *gin.Engine
}

func NewApp() *App {
	return &App{}
}

func (a *App) BuildConfig() *App {
	a.cfg = config.Load()
	return a
}

func (a *App) BuildContainer() *App {
	a.container = container.NewContainer(a.cfg)
	return a
}

func (a *App) BuildHandlers() *App {
	a.handlers = http.NewHandlers(a.container)
	return a
}

func (a *App) BuildRouter() *App {
	gin.SetMode(a.cfg.GinMode)
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
