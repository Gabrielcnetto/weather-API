package clients

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Context = context.Background()
var RedisClient *redis.Client

func Connection() {
	RedisClient = redis.NewClient(&redis.Options{
		DB:   0,
		Addr: "0.0.0.0:6379",
	})
}
