package functions

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/jorgevvs2/dockeryzer/src/utils"
)

func Analyze(name string) {
	if name == "" {
		fmt.Println("Provide a image to analyze")
		return
	}

	ctx := context.Background()
	// Crie um cliente Docker
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	// Obtém a imagem pelo seu nome
	imageInspect, _, err := cli.ImageInspectWithRaw(context.Background(), name)
	if err != nil {
		panic(err)
	}

	utils.PrintImageAnalyzeResults(name, imageInspect, false)
}
