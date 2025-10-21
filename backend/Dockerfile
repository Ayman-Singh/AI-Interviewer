# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main-v4 ./cmd/server

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main-v4 .
COPY --from=builder /app/db ./db

# Give the binary permission to run
RUN chmod +x ./main-v4

EXPOSE 8080

CMD ["./main-v4"]