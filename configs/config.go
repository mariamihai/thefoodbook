package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// NewConfig Order of config loading files: https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func NewConfig() (*Config, error) {
	var config Config

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		log.Println("ENVIRONMENT variable was not set. Will default to 'development'.")
		env = "development"
	}
	config.Environment = env
	log.Printf("Loading environment variables for '%s'.", config.Environment)

	err := godotenv.Load(".env." + env + ".local")
	if err != nil {
		log.Printf("Missing '.env.%s.local'. Using '.env.%s instead'.", env, env)
	}
	if env != "test" {
		log.Println("This is not a test environment. Will be using '.env.local'.")
		err = godotenv.Load(".env.local")
		if err != nil {
			return nil, err
		}
	}
	err = godotenv.Load(".env." + env)
	if err != nil {
		log.Printf("Missing '.env.%s'.", env)
		return nil, err
	}

	// When added .env file
	//err = godotenv.Load()

	config.DBUri = os.Getenv("DB_URI")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPass = os.Getenv("DB_PASS")

	config.Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		config.Port = 3000
	}

	return &config, nil
}

type Config struct {
	Environment string

	Port int

	// DB
	DBUri  string
	DBUser string
	DBPass string
}
