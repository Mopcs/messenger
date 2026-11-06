package config

import "time"

func loadJWT() (JWT, error) {
	accessTTL, err := getEnvAsDuration("JWT_ACCESS_TTL", 15*time.Minute)
	if err != nil {
		return JWT{}, err
	}

	refreshTTL, err := getEnvAsDuration("JWT_REFRESH_TTL", 30*24*time.Hour)
	if err != nil {
		return JWT{}, err
	}

	return JWT{
		AccessSecret: getEnv("JWT_ACCESS_SECRET", ""),
		AccessTTL:    accessTTL,
		RefreshTTL:   refreshTTL,
	}, nil
}
