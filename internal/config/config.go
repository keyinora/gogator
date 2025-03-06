package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// func getConfigFilePath() (string, error) {
// 	home, err := os.UserHomeDir()
// 	if err != nil {
// 		return "", err
// 	}
// 	fullPath := filepath.Join(home, configFileName)
// 	return fullPath, nil
// }

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user home directory: %w", err)
	}

	fullPath := filepath.Join(home, configFileName)

	// Check if the file exists, if not, create it
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err := createDefaultConfigFile(fullPath)
		if err != nil {
			return "", fmt.Errorf("failed to create default config file: %w", err)
		}
	}

	return fullPath, nil
}

func createDefaultConfigFile(path string) error {
	// The username and password should not be included here
	// However as this is a local application and the github is set to private
	// we can leave for now
	defaultConfig := Config{
		DBURL: "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
	}

	data, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal default config: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write default config file: %w", err)
	}

	return nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
