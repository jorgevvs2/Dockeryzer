package functions

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
)

func Analyze(name string) {
	// Crie um cliente Docker
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// Obtém a imagem pelo seu nome
	imageInspect, _, err := cli.ImageInspectWithRaw(context.Background(), name)
	if err != nil {
		panic(err)
	}

	// Obtem a quantidade de camadas (sub-imagens)
	numberOfLayers := len(imageInspect.RootFS.Layers)

	fmt.Printf("A imagem %s possui %d camadas.\n", name, numberOfLayers)
}
