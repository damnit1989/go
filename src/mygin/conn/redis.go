package conn

import "github.com/go-redis/redis/v7"

const (
	REDIS_HOST = "127.0.0.1"
	REDIS_PORT = "6379"
	Password   = ""
	DB         = 0
)

func GetRedis() *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: "",
		DB:       DB,
	})

	return redisCli
}
