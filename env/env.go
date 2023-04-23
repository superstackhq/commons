package env

import "os"

func GetOrDefault(key string, defaultValue string) string {
	val := os.Getenv(key)

	if len(val) == 0 {
		return defaultValue
	}

	return val
}
