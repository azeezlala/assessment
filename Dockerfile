FROM golang:1.23 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates libc6-compat tzdata

# Set working directory
WORKDIR /root/

COPY --from=builder /app/main .

# Set the command to run the binary
CMD ["./main"]

# Expose port 8080
EXPOSE 8080