package analyze

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Analyze() {
	// Crie um cliente Docker
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// Lista todas as imagens Docker
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	// Itera sobre as imagens e inspeciona cada uma
	for _, image := range images {
		inspect, _, err := cli.ImageInspectWithRaw(context.Background(), image.ID)
		if err != nil {
			fmt.Printf("Erro ao inspecionar a imagem %s: %v\n", image.ID, err)
			continue
		}

		// Imprime os detalhes da imagem
		fmt.Printf("Detalhes da Imagem %s:\n", image.ID)
		fmt.Printf("Nome: %s\n", inspect.RepoTags[0])
		fmt.Printf("Tamanho: %d\n", inspect.Size)
		fmt.Printf("Arquitetura: %s\n", inspect.Architecture)
		fmt.Printf("Os: %s\n", inspect.Os)
		fmt.Printf("Autor: %s\n", inspect.Author)
		fmt.Printf("Data de Criação: %s\n", inspect.Created)
		fmt.Printf("Camadas (Sub-Imagens): %d\n", len(inspect.RootFS.Layers))
		fmt.Printf("Configuração: %+v\n", inspect.Config)
		fmt.Println("-------------------------------------")
	}
}
