package functions

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"math"
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

	sizeUnit := "MB"
	sizeInMbs := float32(imageInspect.Size) / float32(math.Pow(10.0, 6))
	sizeInGbs := float32(0.0)

	finalSize := sizeInMbs
	isMoreThanOneGb := sizeInMbs > 1000
	if isMoreThanOneGb {
		sizeUnit = "GB"
		sizeInGbs = sizeInMbs / float32(math.Pow(10.0, 3))
		finalSize = sizeInGbs
	}

	fmt.Printf("%s image details: \n", name)
	fmt.Printf("- Size: %.2f %s\n", finalSize, sizeUnit)
	fmt.Printf("- Layers: %d \n", numberOfLayers)
}
