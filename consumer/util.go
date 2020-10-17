package consumer

import (
	"os"
	"strconv"
)

func readEnvAsInt(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)

	integerEnvValue, err := strconv.Atoi(envValue)

	if err != nil {
		return defaultValue
	}

	return integerEnvValue
}
