package config

import "fmt"

func (c Config) validate() error {

	switch c.App.Env {
	case "development", "production", "test":
	default:
		return fmt.Errorf("APP_ENV must be development, production or test")
	}

	if c.HTTP.Host == "" {
		return fmt.Errorf("HTTP_HOST is required")
	}

	if c.HTTP.Port < 1 || c.HTTP.Port > 65535 {
		return fmt.Errorf("HTTP_PORT must be between 1 and 65535")
	}

	if c.PostgreSQL.DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}

	if c.PostgreSQL.MaxConns <= 0 {
		return fmt.Errorf("POSTGRES_MAX_CONNS must be positive")
	}

	if c.PostgreSQL.MinConns < 0 ||
		c.PostgreSQL.MinConns > c.PostgreSQL.MaxConns {
		return fmt.Errorf("invalid PostgreSQL pool configuration")
	}

	if c.Redis.Addr == "" {
		return fmt.Errorf("REDIS_ADDR is required")
	}

	if c.Redis.DB < 0 {
		return fmt.Errorf("REDIS_DB cannot be negative")
	}

	if len(c.JWT.AccessSecret) < 32 {
		return fmt.Errorf("JWT_ACCESS_SECRET must be at least 32 bytes")
	}

	if c.JWT.AccessTTL <= 0 || c.JWT.RefreshTTL <= 0 {
		return fmt.Errorf("JWT TTL values must be positive")
	}

	return nil
}
