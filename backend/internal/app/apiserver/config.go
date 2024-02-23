package apiserver

type Config struct {
	BindAddr string `toml:"backend_bind_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8081",
		LogLevel: "debug",
	}
}