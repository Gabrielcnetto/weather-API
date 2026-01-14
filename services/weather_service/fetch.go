package weatherservice

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/Gabrielcnetto/weather-API/services/cache"
)

var apiKey = os.Getenv("weather_key")

const apiEndpoint = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"

func FetchWeather(location string) (interface{}, error) {
	var decodedData any
	decodedData, err := cache.GetFromCache(location)
	if decodedData != nil {
		return decodedData, nil
	}
	if err != nil {
		return nil, err
	}
	startDate := time.Now().Format("2006-01-02")
	u, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, location, startDate)
	q := u.Query()
	q.Set("key", apiKey)
	q.Add("unitGroup", "metric")
	q.Add("contentType", "json")
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &decodedData); err != nil {
		return nil, err
	}
	err = cache.SaveCache(location, decodedData)
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}
