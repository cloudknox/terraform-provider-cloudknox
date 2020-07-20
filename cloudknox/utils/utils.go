package utils

import (
	"os"
)

func Truncate(str string, n int, disable bool) string {
	if disable || len(str) < n {
		return str
	} else {
		return str[0:n] + " ... <TRUNCATED>"
	}
}

func CheckIfPathExists(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
