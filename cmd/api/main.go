package main

import "ecommerce_order/app"

func main() {
	app.NewApp().
		BuildConfig().
		BuildContainer().
		BuildHandlers().
		BuildRouter().
		MapWebRoutes().
		Run()
}
