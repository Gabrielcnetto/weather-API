package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	weatherservice "github.com/Gabrielcnetto/weather-API/services/weather_service"
)

func WeatherMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("whithout other verification")
		next.ServeHTTP(w, r)
	})
}

func responseJson(w http.ResponseWriter, r *http.Request, status int, messageKey string, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]interface{}{
		messageKey: message,
	}
	json.NewEncoder(w).Encode(response)
}

func (wi *WeatherInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case http.MethodGet == r.Method:
		response, err := weatherservice.FetchWeather("parobé")
		if err != nil {
			responseJson(w, r, http.StatusBadRequest, "error", err.Error())
			return
		}
		responseJson(w, r, http.StatusOK, "response", response)
		return
	default:
		responseJson(w, r, http.StatusBadRequest, "error", "method not enable")
		return
	}
}
