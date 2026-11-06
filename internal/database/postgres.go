package database

import (
	"Messenger/internal/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context, cfg config.PostgreSQL) (*Postgres, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse PostgreSQL config: %w", err)
	}

	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create PostgreSQL pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	return &Postgres{
		Pool: pool,
	}, nil
}

func (p *Postgres) Ping(ctx context.Context) error {
	if p == nil || p.Pool == nil {
		return fmt.Errorf("PostgreSQL pool is not initialized")
	}

	if err := p.Pool.Ping(ctx); err != nil {
		return fmt.Errorf("ping PostgreSQL: %w", err)
	}

	return nil
}

func (p *Postgres) Close() {
	if p == nil || p.Pool == nil {
		return
	}

	p.Pool.Close()
}
