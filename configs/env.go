package configs

import (
	"os"
	"fmt"
	"gin-test/utils/logs"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "development" {
		env = "dev"
	} else if env == "production" {
		env = "prod"
	}

	filename := fmt.Sprintf(".env.%s", env)

	if err := godotenv.Load(filename); err != nil {
		logs.Error("Error loading env file")
	}
}