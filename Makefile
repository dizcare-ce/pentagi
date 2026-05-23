# Makefile for pentagi project
# Provides common development tasks and shortcuts

.PHONY: all build up down restart logs clean test lint help

# Default target
all: build

# Load environment variables from .env file if it exists
ifneq (,$(wildcard .env))
	include .env
	export
endif

COMPOSE_FILE := docker-compose.yml
DOCKER_COMPOSE := docker compose -f $(COMPOSE_FILE)

## Build all Docker images
build:
	$(DOCKER_COMPOSE) build

## Build without cache
build-no-cache:
	$(DOCKER_COMPOSE) build --no-cache

## Start all services in detached mode
up:
	$(DOCKER_COMPOSE) up -d

## Start all services with logs
up-logs:
	$(DOCKER_COMPOSE) up

## Stop all services
down:
	$(DOCKER_COMPOSE) down

## Stop and remove volumes
down-volumes:
	$(DOCKER_COMPOSE) down -v

## Restart all services
restart: down up

## Show logs for all services (last 100 lines, then follow)
logs:
	$(DOCKER_COMPOSE) logs -f --tail=100

## Show logs for a specific service (usage: make logs-service SERVICE=backend)
logs-service:
	$(DOCKER_COMPOSE) logs -f --tail=100 $(SERVICE)

## Pull latest images
pull:
	$(DOCKER_COMPOSE) pull

## Show status of services
ps:
	$(DOCKER_COMPOSE) ps

## Copy .env.example to .env if .env does not exist
env:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "Created .env from .env.example — please update values before running"; \
	else \
		echo ".env already exists"; \
	fi

## Run Go tests
test:
	go test ./...

## Run Go tests with verbose output
test-verbose:
	go test -v ./...

## Run Go linter
lint:
	golangci-lint run ./...

## Format Go code
fmt:
	gofmt -w .

## Tidy Go modules
tidy:
	go mod tidy

## Clean build artifacts
clean:
	$(DOCKER_COMPOSE) down -v --remove-orphans
	rm -f bin/*

## Build Go binary locally
build-local:
	go build -o bin/pentagi ./cmd/pentagi

## Display help
help:
	@echo "Available targets:"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /' | column -t -s ':'
