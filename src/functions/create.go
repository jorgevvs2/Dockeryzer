package functions

import (
	"fmt"
	"os/exec"
)

func Create(name string) {
	// Caminho para o Dockerfile
	dockerfilePath := "C:\\Users\\jorge\\OneDrive\\Documentos\\local-list\\Dockerfile"

	// Nome da imagem
	imageName := name

	// Construir a imagem Docker usando o Dockerfile
	cmd := exec.Command("docker", "build", "-t", imageName, "-f", dockerfilePath, ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao construir imagem Docker:", err)
		fmt.Println(string(out))
		return
	}

	fmt.Println("Imagem Docker criada com sucesso:", imageName)

	// Criar um contêiner a partir da imagem
	containerName := name
	runCmd := exec.Command("docker", "run", "--name", containerName, "-d", imageName)
	out, err = runCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao criar contêiner Docker:", err)
		fmt.Println(string(out))
		return
	}

	fmt.Println("Contêiner Docker criado com sucesso:", containerName)
}
