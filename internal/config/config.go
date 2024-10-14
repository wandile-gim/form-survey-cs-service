package config

import (
	"github.com/joho/godotenv"
	"log"
	"sync"
)

var (
	once                  sync.Once
	DatabaseConfiguration *Database
)

type Database struct {
	Username        string
	Password        string
	Host            string
	Port            int
	UseDatabase     string
	DefaultDatabase string
}

type Aligo struct {
	ApiKey string
	UserId string
	Sender string
}

func dbConfig() {
	DatabaseConfiguration = &Database{
		Host:            GetEnv("RDB_HOST", "localhost"),
		Port:            GetEnvAsInt("RDB_PORT", 5432),
		Username:        GetEnv("RDB_USER", ""),
		Password:        GetEnv("RDB_PASSWORD", ""),
		UseDatabase:     GetEnv("USE_DATABASE", ""),
		DefaultDatabase: GetEnv("DEFAULT_DATABASE", ""),
	}
}

func SmsAPIConfig() *Aligo {
	return &Aligo{
		ApiKey: GetEnv("ALIGO_API_KEY", ""),
		UserId: GetEnv("ALIGO_USER_ID", ""),
		Sender: GetEnv("ALIGO_SENDER", ""),
	}
}

func tryReadEnvFile() {
	if envFile := GetEnv("CONFIG_FILE", "internal/config/localhost.env"); len(envFile) > 0 {
		err := godotenv.Load(envFile)
		if err != nil {
			log.Printf("could not load .env file: %v", err)
			return
		}
		log.Printf("read env file: %v", envFile)
	}
}

func DefaultSetupFromEnv() {
	tryReadEnvFile()

	once.Do(func() {
		dbConfig()
	})
}
