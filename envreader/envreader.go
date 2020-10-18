package envreader

import (
	"os"
	"strconv"
)

// ReadEnvAsInt reads env variable and converts it to string, if not found or failed returns default value
func ReadEnvAsInt(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)

	integerEnvValue, err := strconv.Atoi(envValue)

	if err != nil {
		return defaultValue
	}

	return integerEnvValue
}
