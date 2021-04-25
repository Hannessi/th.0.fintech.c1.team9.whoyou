package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	MongoServer   []string `json:"mongoServer"`
	ServerPort    string   `json:"serverPort"`
	MongoUsername string   `json:"mongoUsername"`
	MongoPassword string   `json:"mongoPassword"`
	DatabaseName  string   `json:"databaseName"`
}

func LoadConfig(file string) (*Config, error) {
	var config Config

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
