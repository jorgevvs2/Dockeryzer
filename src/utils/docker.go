package utils

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
	"os/exec"
)

func getDockerClient() *client.Client {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Failed to retrieve Docker client")
		fmt.Println(err)
		os.Exit(0)
	}
	cli.NegotiateAPIVersion(ctx)
	return cli
}

func GetDockerImageInspectByIdOrName(idOrName string) types.ImageInspect {
	dockerClient := getDockerClient()

	imageInspect, _, err := dockerClient.ImageInspectWithRaw(context.Background(), idOrName)
	if err != nil {
		fmt.Printf("Failed to retrieve image using the provided name: %s\n", idOrName)
		fmt.Println(err)
		os.Exit(0)
	}

	return imageInspect
}

func ExecDockerBuildCommand(imageName string) *exec.Cmd {
	return exec.Command("docker", "build", "-t", imageName, "-f", "Dockeryzer.Dockerfile", ".")
}
