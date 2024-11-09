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
	if !hasPackageJson() {
		return nil
	}

	readFile, err := os.ReadFile("./package.json")
	if err != nil {
		return nil
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

func GetPackageJsonDependencies() (map[string]string, map[string]string) {
	if !hasPackageJson() {
		return nil, nil
	}

	readFile, err := os.ReadFile("./package.json")
	if err != nil {
		return nil, nil
	}

	var packageJson map[string]interface{}
	err = json.Unmarshal(readFile, &packageJson)
	if err != nil {
		fmt.Println("Error parsing package.json file:", err)
		os.Exit(1)
	}

	dependencies := make(map[string]string)
	devDependencies := make(map[string]string)

	if deps, ok := packageJson["dependencies"].(map[string]interface{}); ok {
		for key, value := range deps {
			if strValue, ok := value.(string); ok {
				dependencies[key] = strValue
			}
		}
	}

	if devDeps, ok := packageJson["devDependencies"].(map[string]interface{}); ok {
		for key, value := range devDeps {
			if strValue, ok := value.(string); ok {
				devDependencies[key] = strValue
			}
		}
	}

	return dependencies, devDependencies
}

func GetRootFiles() []string {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}
