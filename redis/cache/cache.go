package cache

import (
	"CodeLab/redis/cli"
	"context"
	"fmt"
)

type redisCLi struct {
	cli *cli.Redis
}

func Cset(key, value string, ctx context.Context) string {
	redis := cli.NewRedisCLI()
	res, err := redis.Set(ctx, key, value, 0).Result()
	if err != nil {
		fmt.Errorf("failed to set: %v", err)
	}
	return res
}

func Cget(key string, ctx context.Context) string {
	redis := cli.NewRedisCLI()
	res, err := redis.Get(ctx, key).Result()
	if err != nil {
		fmt.Errorf("failed to get: %v", err)
	}
	return res
}

func Cupdate(key, value string, ctx context.Context) string {
	redis := cli.NewRedisCLI()
	res, err := redis.GetSet(ctx, key, value).Result()
	if err != nil {
		fmt.Errorf("failed to get: %v", err)
	}
	return res
}
