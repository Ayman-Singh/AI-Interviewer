# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application - bumped to v7 to break cache
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main-v7 ./cmd/server

# Final stage
FROM alpine:latest

# Use the full mariadb-client to get the required authentication plugins
RUN apk --no-cache add ca-certificates mariadb-client

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main-v7 .
COPY --from=builder /app/db ./db

# Give the binary permission to run
RUN chmod +x ./main-v7

EXPOSE 8080

# Run the init.sql script first, then start the application
# "set -e" will cause the script to exit immediately if the mysql command fails
CMD ["sh", "-c", "set -e; echo '--- Running database migrations ---'; mysql --ssl=0 -h $DB_HOST -P $DB_PORT -u $DB_USER -p$DB_PASSWORD $DB_NAME < ./db/init.sql; echo '--- Migrations complete, starting application ---'; ./main-v7"]

