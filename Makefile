ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Environment configuration
ENV ?= dev
COMPOSE_FILE = docker-compose.$(ENV).yaml

# Check if compose file exists
COMPOSE_FILE_EXISTS = $(shell test -f $(COMPOSE_FILE) && echo yes || echo no)

ifeq ($(COMPOSE_FILE_EXISTS), yes)
	COMPOSE = docker-compose -f $(COMPOSE_FILE)
else
	$(error Compose file $(COMPOSE_FILE) not found! Available: docker-compose.dev.yaml, docker-compose.prod.yaml)
endif

SERVICE_NAME = app

.PHONY: help build up down logs restart rebuild

help:
	@echo "Environment: $(ENV)"
	@echo "Compose file: $(COMPOSE_FILE)"
	@echo ""
	@echo "Commands:"
	@echo "  make build		- Build images"
	@echo "  make up		- Detached containers start"
	@echo "  make up-synced	- Synced containers start"
	@echo "  make down		- Stop containers"
	@echo "  make logs		- Show logs"
	@echo "  make restart	- Restart containers"
	@echo "  make rebuild	- Rebuild and restart"
	@echo "  make config	- Print config"
	@echo "  make status	- Print status"

build:
	$(COMPOSE) build

up-synced:
	$(COMPOSE) up

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

logs:
	$(COMPOSE) logs -f

restart:
	$(COMPOSE) restart

rebuild: down build up

status:
	@echo "=== Containers ==="
	$(COMPOSE) ps
	@echo "==="

# Show current config
config:
	@echo "Environment: $(ENV)"
	@echo "Compose file: $(COMPOSE_FILE)"
	$(COMPOSE) config