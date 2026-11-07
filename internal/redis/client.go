package redis

import (
	"Messenger/internal/config"
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *goredis.Client
}

func NewRedis(ctx context.Context, cfg config.Redis) (*Redis, error) {
	client := goredis.NewClient(&goredis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("ping Redis: %w", err)
	}

	return &Redis{
		Client: client,
	}, nil
}

func (r *Redis) Close() error {
	if r == nil || r.Client == nil {
		return nil
	}

	return r.Client.Close()
}
