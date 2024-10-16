package config

import (
	"github.com/joho/godotenv"
	"log"
	"sync"
)

var (
	once                  sync.Once
	IsDev                 bool
	ClientSecret          string
	EncodeJsonPath        string
	TokenSecretPath       string
	BankAccountSecretPath string
	MessageFormatPath     string
	DatabaseConfiguration *Database
)

func init() {
	DefaultSetupFromEnv()
	isDev()
	loadEnv()
}

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

func isDev() bool {
	IsDev = GetEnvAsBool("DEV", "true")
	return IsDev
}

func loadEnv() {
	ClientSecret = GetEnv("CLIENT_SECRET_PATH", "internal/config/client_secret.json")
	EncodeJsonPath = GetEnv("ENCODE_JSON_PATH", "internal/config/encode.json")
	TokenSecretPath = GetEnv("TOKEN_PATH", "internal/config/token.json")
	BankAccountSecretPath = GetEnv("BANK_ACCOUNT_PATH", "internal/service/config/bank_account.json")
	MessageFormatPath = GetEnv("MESSAGE_FORMAT_PATH", "internal/service/config/message_format.txt")
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
