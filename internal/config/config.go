package config

import (
	"encoding/json"
	"fmt"
	"io"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	db_url	string
	current_user_name	string
}

func Read() Config {
	// reads the JSON file found at ~/.gatorconfig.json and returns a Config struct
}

func SetUser(c Config) {
	// writes the config struct to the JSON file after setting current_user_name

}

func getConfigFilePath() (string, error) {
	var path = "~./" + configFileName
	return path, nil
}

func write(cfg Config) error {
	return nil
}
