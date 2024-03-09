package utils

import (
	"os"
)

func IsViteProject() bool {
	_, err := os.ReadFile("vite.config.js")
	if err != nil {
		return false
	}
	return true
}
