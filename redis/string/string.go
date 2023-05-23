package string

import (
	"CodeLab/redis/cli"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func SetValue(key, value string, ctx context.Context) *redis.StatusCmd {
	cli := cli.NewRedisCLI()
	res := cli.Set(ctx, key, value, time.Second)
	return res
}
func GetSetValue(key, value string, ctx context.Context) *redis.StringCmd {
	cli := cli.NewRedisCLI()
	cmd := cli.GetSet(ctx, key, value)
	return cmd
}
