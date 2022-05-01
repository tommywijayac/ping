package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func Init() (*Config, error) {
	filepath := os.Getenv("GOPATH") + "\\src\\ping\\files\\ping.yaml"
	fmt.Println(filepath)

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("fail to read config file: %w", err)
	}

	cfg := &Config{}
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal config: %w", err)
	}

	return cfg, nil
}
