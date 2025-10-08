# Makefile for AI Interviewer

.PHONY: help install start stop restart logs clean dev-backend dev-frontend test

help:
	@echo "AI Interviewer - Available Commands"
	@echo "===================================="
	@echo "make install     - Install dependencies"
	@echo "make start       - Start all services with Docker"
	@echo "make stop        - Stop all services"
	@echo "make restart     - Restart all services"
	@echo "make logs        - View logs"
	@echo "make clean       - Remove containers and volumes"
	@echo "make dev-backend - Run backend in development mode"
	@echo "make dev-frontend- Run frontend in development mode"
	@echo "make test        - Run tests"

install:
	@echo "Installing dependencies..."
	cd backend && go mod download
	cd frontend && npm install

start:
	@echo "Starting all services..."
	docker-compose up --build -d
	@echo "Services started!"
	@echo "Frontend: http://localhost:3000"
	@echo "Backend: http://localhost:8080"

stop:
	@echo "Stopping all services..."
	docker-compose down

restart:
	@echo "Restarting all services..."
	docker-compose down
	docker-compose up --build -d

logs:
	docker-compose logs -f

clean:
	@echo "Cleaning up..."
	docker-compose down -v
	@echo "Cleanup complete!"

dev-backend:
	@echo "Starting backend in development mode..."
	cd backend && go run cmd/server/main.go

dev-frontend:
	@echo "Starting frontend in development mode..."
	cd frontend && npm run dev

test:
	@echo "Running tests..."
	cd backend && go test ./...
	cd frontend && npm test
