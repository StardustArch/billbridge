FROM golang:1.24

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

# Create a new stage for a smaller final image
FROM alpine:3.21.3

RUN apk --no-cache add ca-certificates
RUN apk --no-cache add libc6-compat

WORKDIR /app

# Copy the binary from builder
COPY --from=0 /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"] 