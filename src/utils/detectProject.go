package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetPackageJsonDependencies() {
	readFile, err := os.ReadFile("package.json")
	if err != nil {
		fmt.Println("Error on read package.json file.", err)
		os.Exit(1)
	}

	var packageJson map[string]any
	err2 := json.Unmarshal(readFile, &packageJson)
	if err2 != nil {
		fmt.Println("Error on parse package.json file.", err2)
		os.Exit(1)
	}

	fmt.Println(packageJson["dependencies"])
}
