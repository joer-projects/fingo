FROM golang:1.22.4-alpine

# Install necessary packages
RUN apk add --no-cache curl git

# Install sqlc
RUN curl -sSL https://github.com/kyleconroy/sqlc/releases/download/v1.26.0/sqlc_1.13.0_linux_amd64.tar.gz | tar -xz -C /usr/local/bin

# Set the working directory
WORKDIR /app

# Copy the configuration and SQL files
COPY . .

# Command to run sqlc generate
CMD ["sqlc", "generate"]
