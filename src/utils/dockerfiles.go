package utils

import (
	"log"
	"os"
)

func getViteDockerfileContent(ignoreComments bool) string {
	if ignoreComments == true {
		return `# --------------------> The build image
FROM node:alpine AS builder

RUN mkdir -p /workspace/app \
    && chown node:node /workspace -R \
    && npm cache clean --force

USER node:node

WORKDIR /workspace/app

COPY --chown=node:node . /workspace/app

RUN npm ci --only=production && npm run build

# --------------------> The production image
FROM node:alpine

COPY --from=builder --chown=node:node /workspace/app/dist /app

USER node

WORKDIR /app

CMD ["npx", "serve", "-p", "3000", "-s", "/app"]
`
	}

	return `# --------------------> The build image
# Use the Node.js image based on Debian Bullseye for the build phase.
# Customization suggestion: You can change the base image to a different version of Node.js, such as 'node:slim' or 'node:alpine', if necessary.
FROM node:alpine AS builder

# Create the /workspace/app directory, set permissions for the 'node' user, and clean npm cache.
RUN mkdir -p /workspace/app \
    && chown node:node /workspace -R \
    && npm cache clean --force

# Set the 'node:node' user to run subsequent commands, ensuring a secure and restricted environment.
USER node:node

# Set the working directory for the application in the build phase.
# Customization suggestion: If your application's working directory is different, you can modify it by changing the value of the WORKDIR variable.
WORKDIR /workspace/app

# Copy all files from the current host directory to the application's working directory in the container.
COPY --chown=node:node . /workspace/app

# Install only production dependencies, optimizing the build process.
RUN npm ci --only=production && npm run build

# --------------------> The production image
# Again, use the Node.js image based on Debian Bullseye for the production phase.
FROM node:alpine

# Copy the compiled files from the build phase to the '/app' directory in the container.
COPY --from=builder --chown=node:node /workspace/app/dist /app

# Set the default user to run subsequent commands.
USER node

# Set the working directory for the application in the production phase.
WORKDIR /app

# Start the static server using the locally installed serve.
# Customization suggestions:
# - The default server port is set to 3000. You can change it as needed by modifying the '-p' argument in the CMD command.
# - If you prefer to use a different server or add more options to the execution command, you can modify the CMD as needed.
CMD ["npx", "serve", "-p", "3000", "-s", "/app"]
`
}

func CreateDockerfileContent(ignoreComments bool) {
	f, err := os.Create("Dockeryzer.Dockerfile")
	if err != nil {
		log.Fatal(err)
	}

	defer DeferCloseFile(f)
	content := getViteDockerfileContent(ignoreComments)

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}
