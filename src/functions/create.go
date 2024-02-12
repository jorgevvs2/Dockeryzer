package functions

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func deferCloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func createDockerfileContent() {
	f, err := os.Create("Dockeryzer.Dockerfile")
	if err != nil {
		log.Fatal(err)
	}

	defer deferCloseFile(f)

	content := "FROM joaovictornsv/portfolio\nEXPOSE 3000"
	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func Create(name string) {
	createDockerfileContent()
	fmt.Println("Dockerfile created")

	if name == "" {
		return
	}

	// Build an image using the created Dockerfile
	cmd := exec.Command("docker", "build", "-t", name, "-f", "Dockeryzer.Dockerfile", ".")
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error on build image:", err)
		fmt.Println(string(out))
		return
	}
	fmt.Printf("Image %s created!\n", name)
}
