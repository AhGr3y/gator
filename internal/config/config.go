package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting file path: %w", err)
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file: %w", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return config, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	newConfig := Config{
		DbURL:           c.DbURL,
		CurrentUserName: c.CurrentUserName,
	}

	err := write(newConfig)
	if err != nil {
		return fmt.Errorf("error writing to config file: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting user home dir: %w", err)
	}

	return homeDir + "/" + configFileName, nil
}

func write(config Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting file path: %w", err)
	}

	err = os.WriteFile(configFilePath, data, 0666)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
