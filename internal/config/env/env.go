package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	LogLevel string
	Token    string
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("runnning the application without a .env file")
	}

	LogLevel = os.Getenv("LOG_LEVEL")

	Token = os.Getenv("TOKEN")
}
