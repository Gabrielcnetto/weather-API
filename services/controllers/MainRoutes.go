package controllers

import (
	"net/http"

	"github.com/Gabrielcnetto/weather-API/services/controllers/handlers"
	"github.com/Gabrielcnetto/weather-API/services/controllers/middlewares"
)

func MainRoutes() {
	mux := http.NewServeMux()
	mux.Handle("/weather", handlers.WeatherMiddleware(&handlers.WeatherInfo{}))
	mainHandler := middlewares.MainMiddleware(mux)
	http.ListenAndServe(":8080", mainHandler)
}
