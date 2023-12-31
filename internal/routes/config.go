package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

type Config struct {
	timeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Cors(next http.Handler) http.Handler { //https://pkg.go.dev/github.com/go-chi/cors#section-readme  catatan
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //boleh semuanya karena pake "*"
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           5,
	}).Handler(next)
}

// set timeout
func (c *Config) SetTimeout(timeInSeconds int) *Config {
	c.timeout = time.Duration(timeInSeconds) * time.Second
	return c
}

// gettimeout
func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}
