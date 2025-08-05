package app

import (
	"context"
	"ecommerce_order/internal/infrastructure/adapters/http"
	"ecommerce_order/internal/infrastructure/config"
	"ecommerce_order/internal/infrastructure/container"
	"log"


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
<<<<<<< HEAD
=======

	log.Println("MONGO_URI =", a.cfg.MongoURI)
    log.Println("MONGO_DATABASE =", a.cfg.MongoDatabase)

>>>>>>> c91927c5bfd9a2e5410ed554a6f898d4a3e4dd7f
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

<<<<<<< HEAD
func (a *App) RunAPI() {
	log.Println("🚀 Starting API on port 8080...")
=======


// 🟡 Este método inicia o consumer em uma goroutine
func (a *App) RunConsumer() *App {
	go func() {
		log.Println("🔁 Iniciando consumidor...")
		if err := a.container.GetOrderConsumer().Consume(context.Background()); err != nil {
			log.Fatalf("❌ Erro ao consumir mensagens: %v", err)
		}
	}()
	return a
}

func (a *App) Run() {
	a.RunConsumer()
>>>>>>> c91927c5bfd9a2e5410ed554a6f898d4a3e4dd7f
	a.router.Run(":8080")
}

func (a *App) RunConsumer() {
	log.Println("🔁 Starting consumer...")
	if err := a.container.GetOrderConsumer().Consume(context.Background()); err != nil {
		log.Fatalf("❌ Error consuming messages: %v", err)
	}
}