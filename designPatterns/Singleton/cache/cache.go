package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
)

type Cache struct {
	cli *redis.Client
}

var (
	RedisCache *Cache
	once       sync.Once
)

func (c *Cache) Set(ctx context.Context, key string, value interface{}) error {
	if err := c.cli.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	res, err := c.cli.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func NewCache(url string) (*Cache, error) {
	once.Do(func() {
		opt, err := redis.ParseURL(url)
		if err != nil {
			return
		}
		cli := redis.NewClient(opt)
		RedisCache = &Cache{
			cli: cli,
		}
	})
	return RedisCache, nil
}

//func NewCache(url string) (*Cache, error) {
//	opt, err := redis.ParseURL(url)
//	if err != nil {
//		return nil, err
//	}
//	cli := redis.NewClient(opt)
//	fmt.Printf("Number of Redis connections: %d\n", cli.PoolStats().TotalConns)
//	return &Cache{
//		cli: cli,
//	}, nil
//}
