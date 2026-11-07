package server

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"Messenger/internal/config"
)

func New(cfg config.HTTP) *http.Server {
	mux := http.NewServeMux()

	return &http.Server{
		Addr:              net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
}
