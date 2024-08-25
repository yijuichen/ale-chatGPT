package environment

import (
	"encoding/json"
	"errors"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	Openai Openai `json:"openai"`
}

type Openai struct {
	ApiKey string `json:"api_key"`
}

func New(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return Config{}, errors.New("API_KEY is required")
	}

	var config Config
	unmarshalErr := json.Unmarshal(data, &config)
	if unmarshalErr != nil {
		return Config{}, unmarshalErr
	}

	config.Openai.ApiKey = apiKey

	return config, nil

}
