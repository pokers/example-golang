package env

import (
	"os"
)

func GetStringEnv(key string) string {
	return os.Getenv(key)
}
