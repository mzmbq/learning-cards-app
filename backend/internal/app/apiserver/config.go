package apiserver

import "os"

type Config struct {
	BindAddr    string `toml:"backend_bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8081",
		LogLevel:    "debug",
		DatabaseURL: os.Getenv("DATABASE_URL"),
		SessionKey:  os.Getenv("SESSION_KEY"),
	}
}
