package main

import (
	"ecommerce_order/app"
	"log"
	"os"
)

const (
	AppModeAPI      = "api"
	AppModeConsumer = "consumer"
)

func main() {
	mode := os.Getenv("APP_MODE")

	appInstance := app.NewApp().
		BuildConfig().
		BuildContainer().
		BuildHandlers().
		BuildRouter().
		MapWebRoutes()

	switch mode {
	case AppModeAPI:
		appInstance.RunAPI()
	case AppModeConsumer:
		appInstance.RunConsumer()
	default:
		log.Fatal("❌ You must define APP_MODE=api or APP_MODE=consumer")
	}
}
