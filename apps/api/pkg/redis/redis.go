package redis

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var Client = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	Password: os.Getenv("REDIS_PASSWORD"),
	DB:       0,
})

func IsKeyMissing(err error) bool {
	return err == redis.Nil
}