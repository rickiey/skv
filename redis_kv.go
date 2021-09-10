package skv

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisKv struct {
	Cli          *redis.Client
	Compress     bool
	CompressType string
}

func SetRedisKv(endpoint, password string) {
	cli := GetRedisCli(endpoint, password)
	DefaultSKV = &RedisKv{Cli: cli}
}

func GetRedisCli(endpoint, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: password,
		PoolSize: 10,
		//Username:     "default",
		MinIdleConns: 1,
		DB:           0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis Connect failed : %v", err.Error())
	}
	return rdb
}

func (r *RedisKv) GetKV(ctx context.Context, key string) ([]byte, error) {
	return r.Cli.Get(ctx, key).Bytes()
}

func (r *RedisKv) PutKV(ctx context.Context, key string, v []byte) error {
	return r.Cli.Set(ctx, key, v, 0).Err()
}
