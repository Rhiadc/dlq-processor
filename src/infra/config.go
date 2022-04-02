package infra

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	MongoURI    string
	MongoDBName string
	Environment string
}

func NewConfig() *Config {
	if os.Getenv("ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env files")
		}
	}

	return &Config{
		MongoURI:    os.Getenv("MONGO_URI"),
		MongoDBName: os.Getenv("MONGO_DBNAME"),
		Environment: os.Getenv("ENVIRONMENT"),
	}
}

func (c Config) IsLocal() bool {
	return c.Environment == "local"
}

func (c Config) IsProduction() bool {
	return c.Environment == "production"
}
