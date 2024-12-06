package lib

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Postgres PostgresConfig
}

func NewConfig() (*Config, error) {
	postgres, err := NewPostgresConfig()
	if err != nil {
		return nil, err
	}
	result := &Config{
		Postgres: *postgres,
	}
	return result, nil
}

func getEnvString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvStringRequired(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("required environment variable %s not found", key)
	}
	return value, nil
}

func getEnvInt(key string, defaultValue int) (int, error) {
	valueAsString := getEnvString(key, "")
	if valueAsString == "" {
		return defaultValue, nil
	}
	valueAsInt, err := strconv.Atoi(valueAsString)
	if err != nil {
		return 0, err
	}
	return valueAsInt, nil
}
