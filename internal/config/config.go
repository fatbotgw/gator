package config

import (
	"encoding/json"
	"fmt"

	// "io"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Database	string	`json:"db_url"`
	CurrentUser	string	`json:"current_user_name"`
}

func Read() Config {
	// reads the JSON file found at ~/.gatorconfig.json and returns a Config struct
	file, err := getConfigFilePath()
	fmt.Printf("Loading config file from: %v\n", file)
	if err != nil {
		fmt.Println("Error on filepath")
		return Config{}
	}
	
	body, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error on file read")
		return Config{}
	}

	var data Config
	if err := json.Unmarshal(body, &data); err != nil {
		return Config{}
	}
	return data
}

func SetUser(c Config) {
	// writes the config struct to the JSON file after setting current_user_name

}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	var path = filepath.Join(homeDir, configFileName)

	return path, nil
}

func write(cfg Config) error {
	return nil
}
