package functions

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func deferCloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func createDockerfileContent() {
	f, err := os.Create("Dockeryzer.Dockerfile")
	if err != nil {
		log.Fatal(err)
	}

	defer deferCloseFile(f)
	content := "# --------------> The build image\n" +
		"FROM node:latest AS builder\n" +
		"RUN mkdir -p /workspace/app && chown node:node /workspace -R\n" +
		"USER node:node\n" +
		"WORKDIR /workspace/app\n" +
		"COPY --chown=node:node . /workspace/app\n" +
		"RUN npm ci --only=production && npm run build\n" +
		"\n" +
		"# --------------> The production image\n" +
		"FROM node@sha256:2f46fd49c767554c089a5eb219115313b72748d8f62f5eccb58ef52bc36db4ad\n" +
		"RUN npm i -g serve\n" +
		"COPY --from=builder  --chown=node:node /workspace/app/dist /app\n" +
		"USER node\n" +
		"ENV NODE_ENV production\n" +
		"WORKDIR /app\n" +
		"ENTRYPOINT [\"serve\", \"-p\", \"3000\",  \"-s\", \"/app\"]\n"

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func Create(name string) {
	createDockerfileContent()
	fmt.Println("Dockerfile created")

	if name == "" {
		return
	}

	// Build an image using the created Dockerfile
	cmd := exec.Command("docker", "build", "-t", name, "-f", "Dockeryzer.Dockerfile", ".")
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error on build image:", err)
		fmt.Println(string(out))
		return
	}
	fmt.Printf("Image %s created!\n", name)
}
