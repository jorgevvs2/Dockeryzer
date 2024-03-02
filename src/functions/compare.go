package functions

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/jorgevvs2/dockeryzer/src/utils"
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

	utils.PrintImageAnalyzeResults(image1, image1Inspect, true, true)
	fmt.Println()
	utils.PrintImageAnalyzeResults(image2, image2Inspect, true, true)
	fmt.Println()

	fmt.Println("Differences:")
	utils.PrintImageCompareLayersResults(image1, image1Inspect, image2, image2Inspect)
	utils.PrintImageCompareSizeResults(image1, image1Inspect, image2, image2Inspect)
	utils.PrintImageCompareNodeJsResults(image1, image1Inspect, image2, image2Inspect)
}
