package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/mzmbq/learning-cards-app/backend/internal/app/apiserver"
)


var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../config.toml", "path to the config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.BindAddr)
	fmt.Println(config.LogLevel)
}