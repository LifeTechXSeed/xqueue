package util

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetBoolEnv(key string, fallback bool) bool {
	value, err := strconv.ParseBool(GetEnv(key, strconv.FormatBool(fallback)))
	if err != nil {
		return fallback
	}
	return value
}

func GetIntEnv(key string, fallback int) int {
	envParam, err := strconv.Atoi(GetEnv(key, string(fallback)))
	if err != nil {
		envParam = fallback
	}

	return envParam
}
