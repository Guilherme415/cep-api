package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token string
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("runnning the application without a .env file")
	}

	Token = os.Getenv("TOKEN")
}
