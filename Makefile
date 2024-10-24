# Define the binary name and Docker image name
SHELL=cmd.exe
BINARY = goTplService
DOCKER_IMAGE = go-tpl-service-image

# Variables
CONFIG_FILE = ./internal/config/config.yaml

# Build the microservice binary for Linux
build:
	@echo Building binary...
	chdir ..\go-tpl-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${BINARY} ./cmd/api
	@echo Done!

# Clean up build artifacts
clean:
	@echo Cleaning up build artifacts...
	@del $(BINARY) || true
	@echo Cleaned up!

# Build Docker image
docker: build
	@echo Building Docker image...
	@docker build -t $(DOCKER_IMAGE) .
	@echo Docker image built: $(DOCKER_IMAGE)

# Start Docker containers using docker-compose
up:
	@echo Starting Docker containers...
	@docker-compose up -d
	@echo Docker containers started!

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build
	@echo Stopping docker images (if running...)
	docker-compose down
	@echo Building (when required) and starting docker images...
	docker-compose up --build -d
	@echo Docker images built and started!

# Stop Docker containers
down:
	@echo Stopping Docker containers...
	@docker-compose down
	@echo Docker containers stopped!

# View logs of the running containers
logs:
	@echo Viewing logs for Docker containers...
	@docker-compose logs -f


# Migration command
# Variables
MIGRATE = migrate
CONFIG = ./internal/config/config.yaml

# Extract configuration values using yq
USERNAME = $(shell yq eval '.database.username' "$(CONFIG)")
PASSWORD = $(shell yq eval '.database.password' "$(CONFIG)")
HOST = $(shell yq eval '.database.host' "$(CONFIG)")
PORT = $(shell yq eval '.database.port' "$(CONFIG)")
DBNAME = $(shell yq eval '.database.name' "$(CONFIG)")
DB_URL = "postgres://$(USERNAME):$(PASSWORD)@$(HOST):$(PORT)/$(DBNAME)?sslmode=disable"

# Create a new migration
create-migration:
	@bash ./db/create_migration.sh

# Migrate Up
migrate-up:
	$(MIGRATE) -path ./db/migrations -database $(DB_URL) up

# Migrate Down (Rollback)
migrate-down:
	$(MIGRATE) -path ./db/migrations -database $(DB_URL) down 1

# Help
help:
	@echo "Usage:"
	@echo "  make create-migration    Create a new migration"
	@echo "  make migrate-up          Run database migrations"
	@echo "  make migrate-down        Rollback database migrations"
	@echo "  make help                Show this help message"
