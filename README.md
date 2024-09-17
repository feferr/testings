
# Go CLI Boilerplate

## Overview

This Go CLI Boilerplate provides a structured starting point for building your own command-line applications. The primary goal is to simplify development by handling common functionality like API requests, configuration management, and containerization, allowing you to focus on implementing business logic.

The boilerplate comes with:
- A well-structured CLI with subcommands.
- Configuration management using Viper and support for environment variables and AWS Secrets Manager.
- Docker support, including debugging with Delve.
- Deployment-ready workflows for Amazon ECS, Fargate or Lambda.

## Table of Contents

- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Commands](#commands)
  - [Info](#info-command)
- [Development](#development)
  - [Running Locally](#running-locally)
  - [Running with Docker](#running-with-docker)
  - [Debugging with Docker](#debugging-with-docker)
- [Deployment](#deployment)
  - [Amazon ECS](#amazon-ecs)
  - [AWS Lambda](#aws-lambda)

---

## Getting Started

### Requirements

- [Go 1.23+](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [AWS CLI](https://aws.amazon.com/cli/)
- [Terraform](https://www.terraform.io/)
- [Viper](https://github.com/spf13/viper) (included in dependencies)

### Installation

1. Clone this repository:

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

### Configuration

Configuration is managed with Viper, which supports YAML files and environment variables.

1. create `.env` or `config.yaml` and set the values

2. Set environment variables:
   ```bash
   export API_URL="https://api.example.com"
   export AWS_REGION="eu-central-1"
   ```

You can also use Amazon Secrets Manager for sensitive configuration data. Integrate it by updating the `config` package.

---

## Commands

### Info Command

The `info` command provides a sample implementation that fetches IP information from a public API.

```bash
go run main.go info
```

You can modify this command to make requests to your own API. See the `business/info.go` file for the core logic.

---

## Development

### Repository Structure
    cmd/                    # Contains command-line interface logic.
        root.go             # Main entry point for CLI and subcommands.
        info.go             # Example command implementation.
    business/               # Business logic, external API calls, etc.
        info.go             # Example business logic for info command.
    config/                 # Configuration handling (Viper and Env variables).
        config.go           # Manages configuration from files or environment.
    .github/                # Contains GitHub Actions workflows.
        workflows/          # Automated deployment YAML files (e.g., ECS, Lambda).
    Dockerfile              # Defines the Docker image for building and running the app.
    docker-compose.yml      # For running the app locally with Docker and Delve.
    go.mod                  # Go module dependencies.
    go.sum                  # Go module checksums.
    main.go                 # Application entry point.


### Running Locally

To run the CLI locally without Docker:

```bash
go run main.go <command>
```

Example:

```bash
go run main.go info
```

---

## Deployment

This boilerplate includes GitHub workflows to automate the deployment of the CLI application to Amazon ECS or Lambda. You can select the desired deployment approach based on your team's requirements.

### Amazon ECS

If you want to deploy to Amazon ECS:

1. Edit the ECS-specific YAML file in the `.github/workflows/deploy-ecs.yml`.
2. Set up your infrastructure using Terraform from the `iac-repo`.
3. Push your changes to the GitHub repository, and the GitHub Actions workflow will handle the build and deploy process.

### AWS Lambda

For deploying as an AWS Lambda function:

1. Edit the Lambda-specific YAML file in the `.github/workflows/deploy_lambda.yml`.
2. Update the necessary infrastructure using Terraform.
3. Push your changes to trigger the Lambda deployment workflow.

---

## Testing

To run tests:

```bash
go test ./...
```

You can also run tests inside Docker by building and running the `Dockerfile.test` if applicable:

```bash
docker build -f Dockerfile.test -t go-cli-app-test .
docker run go-cli-app-test
```

---

## License

This project is licensed under the MIT License.

---

## Acknowledgements

- [Viper](https://github.com/spf13/viper) for configuration management.
- [Cobra](https://github.com/spf13/cobra) for command-line interface development (optional).

---
