FROM golang:1.20-alpine

WORKDIR /app

# Install air for hot reload
RUN go install github.com/cosmtrek/air@latest

# Copy go mod files
COPY go.mod go.sum* ./
RUN go mod download

# Set environment variables
ENV GO111MODULE=on
ENV CGO_ENABLED=0

EXPOSE 8080

# Using air for hot reload
CMD ["air", "-c", ".air.toml"]