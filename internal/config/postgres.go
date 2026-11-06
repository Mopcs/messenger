package config

func loadPostgreSQL() (PostgreSQL, error) {
	maxConns, err := getEnvAsInt32("POSTGRES_MAX_CONNS", 20)
	if err != nil {
		return PostgreSQL{}, err
	}

	minConns, err := getEnvAsInt32("POSTGRES_MIN_CONNS", 5)
	if err != nil {
		return PostgreSQL{}, err
	}

	return PostgreSQL{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		MaxConns:    maxConns,
		MinConns:    minConns,
	}, nil
}
