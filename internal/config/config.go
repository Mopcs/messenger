package config

import (
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
