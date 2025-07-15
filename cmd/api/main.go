package main

import "ecommerce_order/app"
func main() {
	app.NewApp().
		BuildConfig().
		BuildHandlers().
		BuildRouter().
		MapWebRoutes().
		Run()
}