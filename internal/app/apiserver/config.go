package apiserver

type Config struct {
	BindAddr    string   `toml:"backend_bind_addr" envconfig:"BIND_ADDR"`
	LogLevel    string   `toml:"log_level" envconfig:"LOG_LEVEL"`
	DatabaseURL string   `toml:"database_url" envconfig:"DATABASE_URL"`
	SessionKey  string   `toml:"session_key" envconfig:"SESSION_KEY"`
	CORSOrigins []string `toml:"cors_origins" envconfig:"CORS_ORIGINS"`
	RateLimit   int      `toml:"global_rate_limit" envconfig:"GLOBAL_RATE_LIMIT"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		CORSOrigins: nil,
	}
}
