package utils

import (
	"log"
	"os"
)

func getViteDockerfileContent(ignoreComments bool) string {
	if ignoreComments == true {
		return `# --------------------> The build image
FROM node:alpine AS builder

USER node:node

WORKDIR /workspace/app

COPY --chown=node:node . /workspace/app

RUN npm ci --only=production && npm run build && npm cache clean --force

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

# Set the 'node:node' user to run subsequent commands, ensuring a secure and restricted environment.
USER node:node

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

func getGenericDockerfileContent(ignoreComments bool) string {
	if ignoreComments == true {
		return `# --------------------> The build image
FROM node:alpine AS builder

WORKDIR /workspace/app

COPY --chown=node:node package*.json ./

RUN npm ci --only=production && npm cache clean --force

USER node:node

COPY --chown=node:node . .

# --------------------> The production image
FROM node:alpine

WORKDIR /workspace/app

COPY --from=builder --chown=node:node /workspace/app .

USER node

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

# Set the user to run subsequent commands, ensuring a secure environment.
USER node:node

# Copy all files from the local directory to the application's working directory in the container.
COPY --chown=node:node . .

# --------------------> The production image
# Use the Node.js image based on Alpine Linux for the production phase.
FROM node:alpine

# Set the working directory inside the container.
WORKDIR /workspace/app

# Copy files from the build phase to the production environment.
COPY --from=builder --chown=node:node /workspace/app .

# Set the default user to run subsequent commands.
USER node

# Define the command to start the application.
ENTRYPOINT ["npm", "run", "start"]
`
}

func getGenericDockerfileContentWithBuildStep(ignoreComments bool) string {
	if ignoreComments == true {
		return `# --------------------> The build image
FROM node:alpine AS builder

USER node:node

WORKDIR /workspace/app

COPY --chown=node:node . .

RUN npm ci --only=production && npm run build && npm cache clean --force

# --------------------> The production image
FROM node:alpine

WORKDIR /workspace/app

COPY --from=builder --chown=node:node /workspace/app/dist .

USER node

ENTRYPOINT ["npm", "run", "start"]
`
	}

	return `# --------------------> The build image
# Use the Node.js image based on Alpine Linux as the base image for the build phase.
FROM node:alpine AS builder

# Set the user to run subsequent commands, ensuring a secure environment.
USER node:node

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

# Set the default user to run subsequent commands.
USER node

# Define the command to start the application.
ENTRYPOINT ["npm", "run", "start"]
`
}

func getDockerfileContent(ignoreComments bool) string {
	if IsViteProject() {
		return getViteDockerfileContent(ignoreComments)
	}
	if HasBuildCommand() {
		return getGenericDockerfileContentWithBuildStep(ignoreComments)
	}
	return getGenericDockerfileContent(ignoreComments)
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
