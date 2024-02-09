package create

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"io"
	"os"
)

func Create(name string, path string) {
	// Crie um cliente Docker
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// Nome da imagem que você quer criar
	imageName := name

	// Diretório do conteúdo que você quer adicionar à imagem
	contextPath := path

	// Cria um arquivo tar do diretório do conteúdo
	buildContext, err := archive.TarWithOptions(contextPath, &archive.TarOptions{})
	if err != nil {
		panic(err)
	}
	defer buildContext.Close()

	// Opções de construção da imagem
	buildOptions := types.ImageBuildOptions{
		Tags: []string{"node"},
	}

	// Construa a imagem Docker
	buildResponse, err := cli.ImageBuild(context.Background(), buildContext, buildOptions)
	if err != nil {
		panic(err)
	}
	defer buildResponse.Body.Close()

	// Exibir a saída do build
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Imagem criada com sucesso:", imageName)
}
