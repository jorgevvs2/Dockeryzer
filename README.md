# Dockeryzer

Dockeryzer is a tool that helps you to create a Dockerfile for your project.
It is a simple tool that reads the project's dependencies and creates a Dockerfile with the necessary commands to install them.

## Features

- Create a Dockerfile for your project. Dockeryzer uses best practices to create a Dockerfile to optimize the image size.
- Compare two Docker images. Dockeryzer compares two Docker images and shows the differences between them.
- Analyze a Docker image. Dockeryzer shows the details of a Docker image, like the size of the image and the number of layers.

## Benefits
- Save time. Dockeryzer creates a Dockerfile for you, so you don't have to write it manually.
- Optimize your Docker images. Dockeryzer uses best practices to create a Dockerfile, so your Docker images will be smaller and faster.

## Best practices adopted by Dockeryzer
- Use multi-stage builds to reduce the size of the image.
- Use the `COPY` command instead of the `ADD` command.
- Install only the necessary dependencies.
- Use smaller base images like `alpine` or `slim`.

## How to use

Clone this repository:
    
```bash
git clone git@github.com:jorgevvs2/dockeryzer.git
```

Build the project:
```bash
go build . -o dockeryzer
```

Run the project:
```bash
./dockeryzer
```

You can also create an alias to run the project:
```bash
alias dockeryzer="~/path/to/dockeryzer"
```

## Commands

### Create

With the create command you can generate a Dockerfile and create a Docker image (optional).

```bash
dockeryzer create -n imageName
```

Also, you can ignore comments in the Dockerfile with the flag `--ignore-comments` or `-i`.
```bash
dockeryzer create -n imageName -i
```

### Compare

With the compare command you can compare two Docker images.

```bash
dockeryzer compare image1 image2
```

### Analyze

With the analyze command you can analyze a Docker image.

```bash
dockeryzer analyze imageName
```

## How to contribute

If you want to contribute to this project, feel free to open an issue or create a pull request.