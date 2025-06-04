package cfg

import (
	"log"
	"os"
	"strconv"
)

// ForceLoad loads a configuration key from the environment and panics if it is not set.
func ForceLoad(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Required configuration key '" + key + "' is not set")
	}
	return value
}

// Load loads a configuration key from the environment and returns a default value if it is not set.
func Load(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Configuration key '%s' is not set, using default value: %s\n", key, defaultValue)
		// Log the missing key and default value
		return defaultValue
	}
	return value
}

func ForceLoadInt(key string) int {
	value := ForceLoad(key)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic("Required configuration key '" + key + "' is not a valid integer: " + value)
	}
	return intValue
}

func LoadInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Configuration key '%s' is not set, using default value: %d\n", key, defaultValue)
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Configuration key '%s' is not a valid integer: %s, using default value: %d\n", key, value, defaultValue)
		return defaultValue
	}
	return intValue
}
