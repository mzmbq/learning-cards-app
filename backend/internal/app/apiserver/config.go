package apiserver

type Config struct {
	BindAddr    string `toml:"backend_bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8081",
		LogLevel: "debug",
	}
}
