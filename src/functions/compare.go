package functions

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
)

func Compare(image1, image2 string) {
	ctx := context.Background()
	// Crie um cliente Docker
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	// Obtém a imagem pelo seu nome
	image1Inspect, _, errImage1 := cli.ImageInspectWithRaw(context.Background(), image1)
	if errImage1 != nil {
		panic(errImage1)
	}

	image2Inspect, _, errImage2 := cli.ImageInspectWithRaw(context.Background(), image2)
	if errImage2 != nil {
		panic(errImage2)
	}

	// Obtem a quantidade de camadas (sub-imagens)
	numberOfLayers1 := len(image1Inspect.RootFS.Layers)
	numberOfLayers2 := len(image2Inspect.RootFS.Layers)

	if numberOfLayers2 == numberOfLayers1 {
		fmt.Printf("Images have the same number of layers: %d\n", numberOfLayers2)
	} else if numberOfLayers2 > numberOfLayers1 {
		fmt.Printf("Image %s has %d more layers than image %s\n", image2, numberOfLayers2-numberOfLayers1, image1)
	} else {
		fmt.Printf("Image %s has %d more layers than image %s\n", image1, numberOfLayers1-numberOfLayers2, image2)
	}

	if image1Inspect.Size == image2Inspect.Size {
		fmt.Printf("Images have the same size: %d\n Bytes\n", image1Inspect.Size)
	} else if image2Inspect.Size > image1Inspect.Size {
		percent := 100 - (float32(image1Inspect.Size)/float32(image2Inspect.Size))*100
		fmt.Printf("Image %s is %f%% smaller than image %s\n", image1, percent, image2)
	} else {
		percent := 100 - (float32(image2Inspect.Size)/float32(image1Inspect.Size))*100
		fmt.Printf("Image %s is %f%% smaller than image %s\n", image2, percent, image1)
	}
}
