# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application - bumped to v5 to break cache
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main-v5 ./cmd/server

# Final stage
FROM alpine:latest

# CHANGE 1: Add mysql-client to run the SQL script
RUN apk --no-cache add ca-certificates mysql-client

WORKDIR /root/

# Copy the binary from builder
COPY --from-builder /app/main-v5 .
COPY --from-builder /app/db ./db

# Give the binary permission to run
RUN chmod +x ./main-v5

EXPOSE 8080

# CHANGE 2: Run the init.sql script first, then start the application
CMD ["sh", "-c", "mysql -h $DB_HOST -P $DB_PORT -u $DB_USER -p$DB_PASSWORD $DB_NAME < ./db/init.sql && ./main-v5"]
