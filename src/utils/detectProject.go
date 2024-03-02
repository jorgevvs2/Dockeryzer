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

func GetPackageJsonDependencies() map[string]string {
	readFile, err := os.ReadFile("package.json")
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

	return packageJson["dependencies"]
}

func IsViteProject() bool {
	_, err := os.ReadFile("vite.config.js")
	if err != nil {
		return false
	}
	return true
}

func IsExpressProject() bool {
	if !hasPackageJson() {
		return false
	}

	var packageJsonDependencies map[string]string

	packageJsonDependencies = GetPackageJsonDependencies()

	for dependencyName := range packageJsonDependencies {
		if dependencyName == "express" {
			return true
		}
	}
	return false
}
