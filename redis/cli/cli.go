package cli

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
}

func (r *Redis) Get() {

}

func (r *Redis) Set() {

}

func NewRedisCLI() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", //192.168.207.128
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	return rdb
}
