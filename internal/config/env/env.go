package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Host     string
	LogLevel string
	Port     string
	Token    string
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("runnning the application without a .env file")
	}

	LogLevel = os.Getenv("LOG_LEVEL")

	Token = os.Getenv("TOKEN")

	Host = os.Getenv("HOST")
	if Host == "" {
		Host = "localhost"
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = ":8080"
	}
}
