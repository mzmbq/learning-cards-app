package apiserver

type Config struct {
	BindAddr        string          `toml:"backend_bind_addr" envconfig:"BIND_ADDR"`
	LogLevel        string          `toml:"log_level" envconfig:"LOG_LEVEL"`
	DatabaseURL     string          `toml:"database_url" envconfig:"DATABASE_URL"`
	DatabaseTimeout int             `toml:"database_timeout" envconfig:"DATABASE_TIMEOUT"`
	SessionKey      string          `toml:"session_key" envconfig:"SESSION_KEY"`
	CORSOrigins     []string        `toml:"cors_origins" envconfig:"CORS_ORIGINS"`
	GlobaRatelLimit RateLimitConfig `toml:"global_rate_limit" envconfig:"GLOBAL_RATE_LIMIT"`
	UserRateLimit   RateLimitConfig `toml:"user_rate_limit" envconfig:"USER_RATE_LIMIT"`
}

type RateLimitConfig struct {
	Rps     float64 `toml:"rps" envconfig:"RPS"`
	Burst   int     `toml:"burst" envconfig:"BURST"`
	Enabled bool    `toml:"enabled" envconfig:"ENABLED"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		CORSOrigins: nil,
	}
}
