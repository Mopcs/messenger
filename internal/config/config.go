package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type App struct {
	Env  string
	Name string
}

type HTTP struct {
	Host string
	Port int
}

type PostgreSQL struct {
	DatabaseURL string
	MaxConns    int32
	MinConns    int32
}

type Redis struct {
	Addr     string
	Password string
	DB       int
}

type JWT struct {
	AccessSecret string
	AccessTTL    time.Duration
	RefreshTTL   time.Duration
}

type Config struct {
	App        App
	HTTP       HTTP
	PostgreSQL PostgreSQL
	Redis      Redis
	JWT        JWT
}

func Load() (Config, error) {
	cfg := Config{
		App: App{
			Env:  getEnv("APP_ENV", "development"),
			Name: getEnv("APP_NAME", "Messenger"),
		},
		HTTP: HTTP{
			Host: getEnv("HTTP_HOST", "0.0.0.0"),
		},
	}

	port, err := getEnvAsInt("HTTP_PORT", 8080)
	if err != nil {
		return Config{}, err
	}
	cfg.HTTP.Port = port

	cfg.PostgreSQL.DatabaseURL = getEnv("DATABASE_URL", "")

	maxConns, err := getEnvAsInt32("POSTGRES_MAX_CONNS", 20)
	if err != nil {
		return Config{}, err
	}

	minConns, err := getEnvAsInt32("POSTGRES_MIN_CONNS", 5)
	if err != nil {
		return Config{}, err
	}

	cfg.PostgreSQL.MaxConns = maxConns
	cfg.PostgreSQL.MinConns = minConns

	if cfg.PostgreSQL.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}

	if minConns < 0 || maxConns <= 0 || minConns > maxConns {
		return Config{}, fmt.Errorf("invalid PostgreSQL pool configuration")
	}

	cfg.Redis = Redis{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
	}

	db, err := getEnvAsInt("REDIS_DB", 0)
	if err != nil {
		return Config{}, err
	}

	cfg.Redis.DB = db

	if db < 0 {
		return Config{}, fmt.Errorf("REDIS_DB cannot be negative")
	}

	accessTTL, err := getEnvAsDuration("JWT_ACCESS_TTL", 15*time.Minute)
	if err != nil {
		return Config{}, err
	}

	refreshTTL, err := getEnvAsDuration("JWT_REFRESH_TTL", 30*24*time.Hour)
	if err != nil {
		return Config{}, err
	}

	cfg.JWT = JWT{
		AccessSecret: getEnv("JWT_ACCESS_SECRET", ""),
		AccessTTL:    accessTTL,
		RefreshTTL:   refreshTTL,
	}

	if cfg.JWT.AccessSecret == "" {
		return Config{}, fmt.Errorf("JWT_ACCESS_SECRET is required")
	}

	if accessTTL <= 0 || refreshTTL <= 0 {
		return Config{}, fmt.Errorf("JWT TTL values must be positive")
	}

	if cfg.HTTP.Port < 1 || cfg.HTTP.Port > 65535 {
		return Config{}, fmt.Errorf("HTTP_PORT must be between 1 and 65535")
	}

	if cfg.HTTP.Host == "" {
		return Config{}, fmt.Errorf("HTTP_HOST is required")
	}

	if cfg.Redis.Addr == "" {
		return Config{}, fmt.Errorf("REDIS_ADDR is required")
	}

	return cfg, nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func getEnvAsInt(key string, defaultValue int) (int, error) {
	value := getEnv(key, "")

	if value == "" {
		return defaultValue, nil
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value for %s: %w", key, err)
	}

	return result, nil
}

func getEnvAsInt32(key string, defaultValue int32) (int32, error) {
	value := getEnv(key, "")

	if value == "" {
		return defaultValue, nil
	}

	result, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid int32 value for %s: %w", key, err)
	}

	return int32(result), nil
}

func getEnvAsDuration(key string, defaultValue time.Duration) (time.Duration, error) {
	value := getEnv(key, "")

	if value == "" {
		return defaultValue, nil
	}

	result, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("invalid duration value for %s: %w", key, err)
	}

	return result, nil
}
