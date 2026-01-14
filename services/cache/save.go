package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Gabrielcnetto/weather-API/services/clients"
	"github.com/redis/go-redis/v9"
)

func SaveCache(city string, data interface{}) error {
	_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return clients.RedisClient.Set(clients.Context, fmt.Sprintf("city:%v", city), _data, time.Duration(time.Minute*5)).Err()
}

func GetFromCache(city string) (interface{}, error) {
	key := fmt.Sprintf("city:%v", city)

	response, err := clients.RedisClient.Get(clients.Context, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var decodedData interface{}
	if err := json.Unmarshal([]byte(response), &decodedData); err != nil {
		return nil, err
	}

	return decodedData, nil
}
