package handlers

import (
	"encoding/json"
	"net/http"

	weatherservice "github.com/Gabrielcnetto/weather-API/services/weather_service"
)

func responseJson(w http.ResponseWriter, r *http.Request, status int, messageKey string, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]interface{}{
		messageKey: message,
	}
	json.NewEncoder(w).Encode(response)
}
func WeatherMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		location := r.Header.Get("city")
		if location == "" {
			responseJson(w, r, http.StatusBadRequest, "error", "Location not found")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (wi *WeatherInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case http.MethodGet == r.Method:
		location := r.Header.Get("city")
		response, err := weatherservice.FetchWeather(location)
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
