package common

import (
	"os"
)

func getENVValue(key, defaultValue string) (v string) {
	var found bool
	v, found = os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return v
}
