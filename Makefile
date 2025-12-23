# GLOBALS
API_DIR=apps/api
COMPOSE=docker compose

# CODE GENERATION

.PHONY: sqlc
sqlc:
	@echo "▶ Generating sqlc code..."
	cd $(API_DIR) && sqlc generate

# DATABASE (GOOSE)

.PHONY: migrate-up
migrate-up:
	@echo "▶ Running DB migrations (up)..."
	cd $(API_DIR) && goose -env ../../.env up

.PHONY: migrate-down
migrate-down:
	cd $(API_DIR) && goose -env ../../.env down

.PHONY: migrate-status
migrate-status:
	cd $(API_DIR) && goose -env ../../.env status

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "❌ Usage: make migrate-create name=your_migration_name"; \
		exit 1; \
	fi
	cd $(API_DIR) && goose create $(name) sql

# DOCKER

.PHONY: docker-build
docker-build:
	$(COMPOSE) build

.PHONY: docker-up
docker-up:
	$(COMPOSE) up

.PHONY: docker-up-build
docker-up-build:
	$(COMPOSE) up --build

.PHONY: docker-down
docker-down:
	$(COMPOSE) down

.PHONY: docker-reset
docker-reset:
	$(COMPOSE) down -v
	$(COMPOSE) up --build

# LOCAL DEV

.PHONY: api-run
api-run:
	cd $(API_DIR) && go run cmd/server/main.go

.PHONY: tidy
tidy:
	cd $(API_DIR) && go mod tidy

# HELP

.PHONY: help
help:
	@echo ""
	@echo "Available commands:"
	@echo "  make sqlc                → Generate sqlc code"
	@echo "  make migrate-up          → Run DB migrations"
	@echo "  make migrate-down        → Rollback last migration"
	@echo "  make migrate-status      → Migration status"
	@echo "  make migrate-create name=xyz"
	@echo "  make docker-up-build     → Docker up with build"
	@echo "  make docker-reset        → Full reset (⚠ deletes DB)"
	@echo "  make api-run             → Run API locally"
	@echo ""
