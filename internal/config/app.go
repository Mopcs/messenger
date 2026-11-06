package config

func loadApp() App {
	return App{
		Env:  getEnv("APP_ENV", "development"),
		Name: getEnv("APP_NAME", "Messenger"),
	}
}

func loadHTTP() (HTTP, error) {
	port, err := getEnvAsInt("HTTP_PORT", 8080)
	if err != nil {
		return HTTP{}, err
	}

	return HTTP{
		Host: getEnv("HTTP_HOST", "0.0.0.0"),
		Port: port,
	}, nil
}
