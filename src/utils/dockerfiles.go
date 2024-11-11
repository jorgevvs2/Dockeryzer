package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jorgevvs2/dockeryzer/src/config"
	"github.com/sashabaranov/go-openai"
)

type ProjectInfo struct {
	HasPackageJson  bool              `json:"hasPackageJson"`
	RootFiles       []string          `json:"rootFiles"`
	Scripts         map[string]string `json:"scripts,omitempty"`
	Dependencies    map[string]string `json:"dependencies,omitempty"`
	DevDependencies map[string]string `json:"devDependencies,omitempty"`
}

func getViteDockerfileContent(ignoreComments bool) string {
	if ignoreComments == true {
		return `# --------------------> The build image
FROM node:alpine AS builder

WORKDIR /workspace/app

COPY --chown=node:node . /workspace/app

RUN npm ci --only=production && npm run build && npm cache clean --force

# --------------------> The production image
FROM node:alpine

COPY --from=builder --chown=node:node /workspace/app/dist /app

WORKDIR /app

CMD ["npx", "serve", "-p", "3000", "-s", "/app"]
`
	}

	return `# --------------------> The build image
# Use the Node.js image based on Debian Bullseye for the build phase.
# Customization suggestion: You can change the base image to a different version of Node.js, such as 'node:slim' or 'node:alpine', if necessary.
FROM node:alpine AS builder

# Set the working directory for the application in the build phase.
# Customization suggestion: If your application's working directory is different, you can modify it by changing the value of the WORKDIR variable.
WORKDIR /workspace/app

# Copy all files from the current host directory to the application's working directory in the container.
COPY --chown=node:node . /workspace/app

# Install only production dependencies, optimizing the build process, and clean npm cache.
RUN npm ci --only=production && npm run build && npm cache clean --force

# --------------------> The production image
# Again, use the Node.js image based on Debian Bullseye for the production phase.
FROM node:alpine

# Copy the compiled files from the build phase to the '/app' directory in the container.
COPY --from=builder --chown=node:node /workspace/app/dist /app

# Set the working directory for the application in the production phase.
WORKDIR /app

# Start the static server using the locally installed serve.
# Customization suggestions:
# - The default server port is set to 3000. You can change it as needed by modifying the '-p' argument in the CMD command.
# - If you prefer to use a different server or add more options to the execution command, you can modify the CMD as needed.
CMD ["npx", "serve", "-p", "3000", "-s", "/app"]
`
}

func getGenericDockerfileContent(ignoreComments bool) string {
	if ignoreComments == true {
		return `# --------------------> The build image
FROM node:alpine AS builder

WORKDIR /workspace/app

COPY --chown=node:node package*.json ./

RUN npm ci --only=production && npm cache clean --force

COPY --chown=node:node . .

# --------------------> The production image
FROM node:alpine

WORKDIR /workspace/app

COPY --from=builder --chown=node:node /workspace/app .

ENTRYPOINT ["npm", "run", "start"]
`
	}

	return `# --------------------> The build image
# Use the Node.js image based on Alpine Linux as the base image for the build phase.
FROM node:alpine AS builder

# Set the working directory inside the container.
WORKDIR /workspace/app

# Copy package files (package.json and package-lock.json) to the container.
COPY --chown=node:node package*.json ./

# Install only production dependencies and clean npm cache to optimize the build process.
RUN npm ci --only=production && npm cache clean --force

# Copy all files from the local directory to the application's working directory in the container.
COPY --chown=node:node . .

# --------------------> The production image
# Use the Node.js image based on Alpine Linux for the production phase.
FROM node:alpine

# Set the working directory inside the container.
WORKDIR /workspace/app

# Copy files from the build phase to the production environment.
COPY --from=builder --chown=node:node /workspace/app .

# Define the command to start the application.
ENTRYPOINT ["npm", "run", "start"]
`
}

func getGenericDockerfileContentWithBuildStep(ignoreComments bool) string {
	if ignoreComments == true {
		return `# --------------------> The build image
FROM node:alpine AS builder

WORKDIR /workspace/app

COPY --chown=node:node . .

RUN npm ci --only=production && npm run build && npm cache clean --force

# --------------------> The production image
FROM node:alpine

WORKDIR /workspace/app

COPY --from=builder --chown=node:node /workspace/app/dist .

ENTRYPOINT ["npm", "run", "start"]
`
	}

	return `# --------------------> The build image
# Use the Node.js image based on Alpine Linux as the base image for the build phase.
FROM node:alpine AS builder

# Set the working directory inside the container.
WORKDIR /workspace/app

# Copy all files from the local directory to the application's working directory in the container.
COPY --chown=node:node . .

# Install only production dependencies, build the project and clean npm cache to optimize the build process.
RUN npm ci --only=production && npm run build && npm cache clean --force

# --------------------> The production image
# Use the Node.js image based on Alpine Linux for the production phase.
FROM node:alpine

# Set the working directory inside the container.
WORKDIR /workspace/app

# Copy build files from the build phase to the production environment.
COPY --from=builder --chown=node:node /workspace/app/dist .

# Define the command to start the application.
ENTRYPOINT ["npm", "run", "start"]
`
}

func generateAIPrompt(info ProjectInfo, ignoreComments bool) string {
	// Convert project info to a concise JSON string
	infoJson, _ := json.MarshalIndent(info, "", "  ")

	basePrompt := `Generate a production-ready optimized Dockerfile for a Node.js project with the following characteristics:
%s

Technical requirements:
- If you detect a Next.js project, please use the appropriate Node.js base image. Otherwise, use the latest LTS version, i.e. "node:alpine"
- The Dockerfile must be optimized for production use
- If you detect a frontend project, feel free to use other tools to serve the application, e.g. npx serve, nginx, etc. Make sure to import a valid base image for the tool (e.g. node:alpine for npx serve)
- Use multi-stage builds to optimize the final image size
- Try to keep the number of layers as low as possible
- Try to keep the final image size as small as possible
- Follow security best practices
- Include only necessary files
- Make sure the application starts correctly
- If you detect a different package manager (e.g. Yarn, pnpm), make sure to install it on the production image
- Install dev dependencies if necessary
- Make sure to copy all necessary files to the production image
- At the end of the Dockerfile, add a comment with the "docker run" example command to start the application.

Formatting requirements:
- Return ONLY the raw Dockerfile content without any markdown formatting, code blocks, or explanations
- Start directly with the FROM instruction or the comment block
- Do not include any markdown backticks or formatting
%s

Remember:
Respond with only the raw Dockerfile content, starting with FROM (or the comment block) and no other text or formatting.`

	commentInstruction := ""
	if ignoreComments {
		commentInstruction = "- Do not include any comments in the Dockerfile"
	} else {
		commentInstruction = "- Each instruction must be preceded by a comment explaining its purpose\n- Comments must be on their own lines, above their related instructions"
	}

	return fmt.Sprintf(basePrompt, string(infoJson), commentInstruction)
}

func getDockerfileContent(ignoreComments bool) string {
	// Gather project information
	info := ProjectInfo{
		HasPackageJson: hasPackageJson(),
		RootFiles:      GetRootFiles(),
	}

	// Only include package.json related info if it exists
	if info.HasPackageJson {
		info.Scripts = GetPackageJsonScripts()
		info.Dependencies, info.DevDependencies = GetPackageJsonDependencies()
	}

	// Generate AI prompt
	prompt := generateAIPrompt(info, ignoreComments)

	// Use the embedded API key
	apiKey := config.APIKey
	if apiKey == "" {
		log.Fatal("API key not set in binary. Please rebuild with -ldflags")
	}

	// Show loading message
	fmt.Println("🤖 AI is analyzing your project and generating a Dockerfile...")

	// Initialize OpenAI client
	client := openai.NewClient(apiKey)

	// Create chat completion request
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a Docker expert. Respond only with Dockerfile content, no explanations.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0.2,
		},
	)

	if err != nil {
		fmt.Println("❌ Error generating Dockerfile with AI, falling back to default logic...")
		if IsViteProject() {
			return getViteDockerfileContent(ignoreComments)
		}
		if HasBuildCommand() {
			return getGenericDockerfileContentWithBuildStep(ignoreComments)
		}
		return getGenericDockerfileContent(ignoreComments)
	}

	fmt.Println("✅ Dockerfile generated successfully!")

	// Get the response content and clean up any potential markdown
	dockerfile := resp.Choices[0].Message.Content
	dockerfile = strings.TrimSpace(dockerfile)

	// Remove any Markdown code block indicators if present
	dockerfile = strings.TrimPrefix(dockerfile, "```dockerfile")
	dockerfile = strings.TrimPrefix(dockerfile, "```")
	dockerfile = strings.TrimSuffix(dockerfile, "```")
	dockerfile = strings.TrimSpace(dockerfile)

	return dockerfile
}

func CreateDockerfileContent(ignoreComments bool) {
	f, err := os.Create("Dockeryzer.Dockerfile")
	if err != nil {
		log.Fatal(err)
	}

	defer DeferCloseFile(f)
	content := getDockerfileContent(ignoreComments)

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}
