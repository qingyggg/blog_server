package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvInit() {
	mode := os.Getenv("FOO_ENV")
	//load common env
	LoadEnvFile(".env")
	if "production" == mode {
		LoadEnvFile("prod" + ".env")
	} else {
		LoadEnvFile("dev" + ".env")
	}
}

// LoadEnvFile please put env file in /envs directory before using this func
func LoadEnvFile(fileName string) {
	envFile, err := os.Open("envs/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	envMap, err := godotenv.Parse(envFile)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range envMap {
		err := os.Setenv(k, v)
		if err != nil {
			log.Fatal(err)
		}
	}
}
