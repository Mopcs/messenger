package config

import "fmt"

func Load() (Config, error) {
	var cfg Config
	var err error

	cfg.App = loadApp()

	cfg.HTTP, err = loadHTTP()
	if err != nil {
		return Config{}, fmt.Errorf("load HTTP config: %w", err)
	}

	cfg.PostgreSQL, err = loadPostgreSQL()
	if err != nil {
		return Config{}, fmt.Errorf("load PostgreSQL config: %w", err)
	}

	cfg.Redis, err = loadRedis()
	if err != nil {
		return Config{}, fmt.Errorf("load Redis config: %w", err)
	}

	cfg.JWT, err = loadJWT()
	if err != nil {
		return Config{}, fmt.Errorf("load JWT config: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return Config{}, fmt.Errorf("validate config: %w", err)
	}

	return cfg, nil
}
