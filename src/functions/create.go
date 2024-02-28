package functions

import (
	"bufio"
	"fmt"
	"github.com/jorgevvs2/dockeryzer/src/utils"
	"os/exec"
)

func Create(name string) {
	utils.CreateDockerfileContent()
	utils.CreateDockerignoreContent()

	successOut := utils.GetSuccessOutput()
	infoOut := utils.GetInfoOutput()
	errorOut := utils.GetErrorOutput()

	fmt.Println("New files:")
	successOut.Println("\tDockeryzer.Dockerfile\n\t.dockerignore")

	if name == "" {
		fmt.Println("\nTo build your image, run one of the following commands::")
		fmt.Println("- To specify a name for the image:")
		infoOut.Println("\tdocker build -t <image-name> -f Dockeryzer.Dockerfile .")
		fmt.Println("- To build without specifying a name:")
		infoOut.Println("\tdocker build -f Dockeryzer.Dockerfile .")
		return
	}

	infoOut.Printf("\nBuilding your image %s...\n", name)
	cmd := exec.Command("docker", "build", "-t", name, "-f", "Dockeryzer.Dockerfile", ".")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error on create pipe to handle stdout", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error on create pipe to handle stderr:", err)
		return
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error on start command:", err)
		return
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
		errorOut.Println("Error on waiting command finish:", err)
		return
	}
}
