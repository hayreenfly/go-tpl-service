FROM alpine:latest

# Create app directory
RUN mkdir /app

# Set the working directory
WORKDIR /app

# Copy the compiled binary and config.yaml into the app directory
COPY goTplService /app/
COPY internal/config/config.yaml /app/internal/config/config.yaml

# Ensure the binary has execution permission
RUN chmod +x /app/goTplService

# Command to run the application
CMD ["/app/goTplService"]
