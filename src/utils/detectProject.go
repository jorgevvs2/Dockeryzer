package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func hasPackageJson() bool {
	_, err := os.ReadFile("package.json")
	if err != nil {
		return false
	}
	return true
}

func GetPackageJsonScripts() map[string]string {
	readFile, err := os.ReadFile("./package.json")
	if err != nil {
		fmt.Println("Error on read package.json file.", err)
		os.Exit(1)
	}

	var packageJson map[string]map[string]string
	err2 := json.Unmarshal(readFile, &packageJson)
	if err2 != nil {
		fmt.Println("Error on parse package.json file.", err2)
		os.Exit(1)
	}

	return packageJson["scripts"]
}

func IsViteProject() bool {
	_, err := os.ReadFile("vite.config.js")
	if err != nil {
		return false
	}
	return true
}

func HasBuildCommand() bool {
	if !hasPackageJson() {
		return false
	}

	var packageJsonScripts map[string]string

	packageJsonScripts = GetPackageJsonScripts()
	if packageJsonScripts == nil {
		fmt.Println("Error on get package.json scripts.")
		os.Exit(1)
	}

	for script := range packageJsonScripts {
		if script == "build" {
			return true
		}
	}
	return false
}
