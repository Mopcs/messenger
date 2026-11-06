package config

func loadRedis() (Redis, error) {
	db, err := getEnvAsInt("REDIS_DB", 0)
	if err != nil {
		return Redis{}, err
	}

	return Redis{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       db,
	}, nil
}
