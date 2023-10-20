# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the current working directory inside the container to /app
WORKDIR /app

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 80

HEALTHCHECK --interval=10s --timeout=3s --retries=2 CMD wget --quiet --tries=1 --spider http://localhost:80/health || exit 1

#Command to run the executable
CMD ["./main"]
