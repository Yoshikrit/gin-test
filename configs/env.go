package configs

import (
	"os"
	"fmt"
	"gin-test/utils/logs"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	envName := ""
	if env == "development" {
		envName = "dev"
	} else if env == "production" {
		envName = "prod"
	} else {
		logs.Error("Error no environment") 
	}

	filename := fmt.Sprintf(".env.%s", envName)

	if err := godotenv.Load(filename); err != nil {
		logs.Error("Error loading env file")
		panic(err)
	}
}