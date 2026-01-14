package main

import (
	"github.com/Gabrielcnetto/weather-API/services/clients"
	"github.com/Gabrielcnetto/weather-API/services/controllers"
)

func main() {
	clients.Connection()
	controllers.MainRoutes()
}
