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
		fmt.Println("Error reading package.json file:", err)
		os.Exit(1)
	}

	var packageJson map[string]interface{}
	err = json.Unmarshal(readFile, &packageJson)
	if err != nil {
		fmt.Println("Error parsing package.json file:", err)
		os.Exit(1)
	}

	// Safely access the "scripts" field and assert its type
	scripts := make(map[string]string)
	if scriptsSection, ok := packageJson["scripts"].(map[string]interface{}); ok {
		for key, value := range scriptsSection {
			if strValue, ok := value.(string); ok {
				scripts[key] = strValue
			} else {
				fmt.Printf("Warning: script %q is not a string and was skipped\n", key)
			}
		}
	} else {
		fmt.Println("Error: 'scripts' section is not a valid map[string]string")
		os.Exit(1)
	}

	return scripts
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
