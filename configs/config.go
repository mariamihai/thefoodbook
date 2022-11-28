package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func InitializeConfiguration() Config {
	return InitializeConfigurationForFilename("../local.env")
}

func InitializeConfigurationForFilename(filename string) Config {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error encountered when loading the configuration: %s", err)
	}

	var config Config
	config.Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("couldn't get configured PORT: %s", err)
	}

	return config
}

type Config struct {
	Port int

	// TODO - connect to neo4j
	Uri      string
	Username string
	Password string
}
