package functions

import (
	"bufio"
	"fmt"
	"github.com/jorgevvs2/dockeryzer/src/utils"
	"os"
	"os/exec"
)

func Create(imageName string, ignoreComments bool) {
	utils.CreateDockerfileContent(ignoreComments)
	utils.CreateDockerignoreContent()

	fmt.Println("New files:")
	utils.SuccessPrintf("\tDockeryzer.Dockerfile\n\t.dockerignore\n")

	if imageName == "" {
		fmt.Println("\nTo build your image, run one of the following commands::")
		fmt.Println("- To specify a imageName for the image:")
		utils.InfoPrintf("\tdocker build -t <image-imageName> -f Dockeryzer.Dockerfile .\n")
		fmt.Println("- To build without specifying a imageName:")
		utils.InfoPrintf("\tdocker build -f Dockeryzer.Dockerfile .\n")
		return
	}

	utils.InfoPrintf("\nBuilding your image %s...\n", imageName)
	cmd := exec.Command("docker", "build", "-t", imageName, "-f", "Dockeryzer.Dockerfile", ".")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		utils.ErrorPrintf("Error on create pipe to handle stdout: %s\n", err)
		os.Exit(0)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		utils.ErrorPrintf("Error on create pipe to handle stderr: %s\n", err)
		os.Exit(0)
	}

	err = cmd.Start()
	if err != nil {
		utils.ErrorPrintf("Error on start command: %s\n", err)
		os.Exit(0)
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	err = cmd.Wait()
	if err != nil {
		utils.ErrorPrintf("Error on waiting command finish: %s\n", err)
		os.Exit(0)
	}
}
