package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Client *redis.Client

func Init(addr string) {
	Client = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
