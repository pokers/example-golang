package util

import (
	"os"
)

func GetEnvString(key string) string {
	return os.Getenv(key)
}
