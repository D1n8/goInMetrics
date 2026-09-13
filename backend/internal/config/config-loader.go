package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server 		ServerConfig 		`toml:"server"`
	Statistics 	StatisticsConfig 	`toml:"statistics"`
}

type ServerConfig struct {
	Port		int 			`toml:"PORT"`
}

type StatisticsConfig struct {
	UpdateInterval	int 			`toml:"UPDATE_INTERVAL"`
}


func ConfigLoader() (*Config, error){
	cfgPath, err := ConfigPath()
	if err != nil {
		return nil, err
	}

	// Reading bytes from config file
	byteData, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var cfg Config

	toml.Unmarshal(byteData, &cfg)
	return &cfg, nil
}
