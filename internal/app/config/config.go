package config

import (
	"os"
)

func GetDef(key string, defaultValue string) (val string) {
	if val = os.Getenv(key); len(val) > 0 {
		return
	} else {
		val = defaultValue
	}
	return
}
