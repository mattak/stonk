package util

import (
	"log"
	"os"
)

func LoadEnv(key string) string {
	v := os.Getenv(key)
	if len(v) <= 0 {
		log.Fatalln("ERROR: missing environment variable: ", key)
	}
	return v
}
