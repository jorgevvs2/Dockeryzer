package functions

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
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

	// Obtem a quantidade de camadas (sub-imagens)
	numberOfLayers := len(imageInspect.RootFS.Layers)

	fmt.Printf("%s image details: \n", name)
	fmt.Printf("- Size: %d Bytes\n", imageInspect.Size)
	fmt.Printf("- Layers: %d \n", numberOfLayers)
}
