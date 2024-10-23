# Define the binary name and Docker image name
SHELL=cmd.exe
BINARY = goTplService
DOCKER_IMAGE = go-tpl-service-image

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
