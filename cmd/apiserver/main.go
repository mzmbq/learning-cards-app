package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./config.toml", "path to the config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	// First load the configuration from the config file
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	// Then load the configuration from the environment variables (overrides the config file)
	err = envconfig.Process("", config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
