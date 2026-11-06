package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func getEnvAsInt(key string, defaultValue int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", key, err)
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
		return 0, fmt.Errorf("parse %s: %w", key, err)
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
		return 0, fmt.Errorf("parse %s: %w", key, err)
	}

	return result, nil
}
