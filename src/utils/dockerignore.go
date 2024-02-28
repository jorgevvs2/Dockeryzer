package utils

import (
	"log"
	"os"
)

func getDockerignoreContent() string {
	dockerignoreContent := `# Versioning configuration files (git, svn, etc.)
.git
.svn

# Build files
node_modules/
npm-debug.log
yarn-error.log
.pnpm-debug.log

# Environment files and directories
.env
.env.local
.env.*.local
npm-debug.log*
yarn-debug.log*
yarn-error.log*
.pnpm-debug.log*

# Operating system-specific local dependencies
.DS_Store
Thumbs.db

# Editor files
.vscode/
.idea/
*.suo
*.ntvs*
*.njsproj
*.sln
*.swp

# Package dependencies and cache files
.npm
.pnpm
.pnp
node_modules
dist

# Compilation files
*.log
*.log*
*.swp
*.tgz
*.zip
*.gz`

	return dockerignoreContent
}

func CreateDockerignoreContent() {
	f, err := os.Create(".dockerignore")
	if err != nil {
		log.Fatal(err)
	}

	defer DeferCloseFile(f)
	content := getDockerignoreContent()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}
