package server

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
