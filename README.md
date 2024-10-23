# Go Microservice Template

This is a basic template for a Go microservice that demonstrates a simple API with example routes and Swagger documentation.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Database Migration](#database-migration)
- [Swagger Documentation](#swagger-documentation)

## Requirements

- Go 1.18 or higher
- Git
- postgres

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/hayreenfly/go-tpl-template.git new-repo-name
   cd new-repo-name
   ```

2. Install Dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `config.yaml` file by copying the example configuration:

   ```bash
   cp ./internal/config/config_example.yaml ./internal/config/config_dev.yaml
   ```

4. Edit the `config.yaml` file to set your specific configurations as needed.

## Usage

1. **Start the Server**:
   Run the following command to start the service:

   ```bash
   go run cmd/api/main.go
   ```

## Database Migration

### Running Migrations

Create a new migration

Create a new migration with the following command:

```bash
make create-migration
```

**Apply migrations**

To apply the migrations, use:

```bash
make migration-up
```

**Rollback migrations**

To rollback the migrations, use:

```bash
make migration-down
```

## Swagger Documentation

This project utilizes Swagger for API documentation. Follow the steps below to generate and view the Swagger documentation.

```bash
swag init -g ./cmd/api/main.go -o ./docs
```
