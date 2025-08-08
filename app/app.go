package app

import (
	"context"
	"ecommerce_order/internal/infrastructure/adapters/http"
	"ecommerce_order/internal/infrastructure/config"
	"ecommerce_order/internal/infrastructure/container"
	"log"
	"os"
	"os/signal"
	"syscall"

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

func (a *App) RunAPI() {
	a.router.Run(":8080")
}

func (a *App) RunConsumer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		log.Println("[app] [msg:shutdown signal received. Cancelling context...]")
		cancel()
	}()

	if err := a.container.GetOrderConsumer().Consume(ctx); err != nil {
		log.Printf("[feature:consumer] [msg:error consuming messages] [error: %v]", err)
	}
}