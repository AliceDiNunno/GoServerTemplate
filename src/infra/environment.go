package infra

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Unable to load env file.")
	}
}

func requireEnvVariable(name string) string {
	envVariable, exists := os.LookupEnv(name)

	if exists == false {
		log.Fatalln("Env variable: ", name, " is not found")
	}

	if strings.Contains(name, "SECRET") || strings.Contains(name, "TOKEN") || strings.Contains(name, "PASSWORD") {
		log.Println("[", name, "] = ", "****")
	} else {
		log.Println("[", name, "] = ", envVariable)
	}

	return envVariable
}
