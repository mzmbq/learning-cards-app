package apiserver

type Config struct {
	BindAddr    string `toml:"backend_bind_addr" envconfig:"BIND_ADDR"`
	LogLevel    string `toml:"log_level" envconfig:"LOG_LEVEL"`
	DatabaseURL string `toml:"database_url" envconfig:"DATABASE_URL"`
	SessionKey  string `toml:"session_key" envconfig:"SESSION_KEY"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
