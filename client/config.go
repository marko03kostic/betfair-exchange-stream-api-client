package client

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppKey  string `json:"appKey"`
	Session string `json:"session"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}