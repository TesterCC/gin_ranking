package cache

import (
	"context"
	"gin_ranking/config"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword, // no password set
		DB:       config.RedisDb,       // use default DB
	})
	Rctx = context.Background()
}

// for ZAdd use
func Zscore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}