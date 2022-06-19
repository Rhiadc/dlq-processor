package infra

import (
	"os"

	"github.com/joho/godotenv"
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
			return &Config{
				MongoURI:    "mongodb://root:s3cr3t@localhost:27017",
				MongoDBName: "dlq",
				Environment: "",
			}
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
