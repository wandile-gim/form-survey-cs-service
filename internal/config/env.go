package config

import (
	"os"
	"strconv"
)

func GetEnv(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

func GetEnvAsInt(key string, defaultVal int) int {
	strVal := GetEnv(key, "")

	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}
	return defaultVal
}

func GetEnvAsBool(key string, defaultVal string) bool {
	strVal := GetEnv(key, defaultVal)

	if strVal == "false" {
		return false
	}
	return true
}
