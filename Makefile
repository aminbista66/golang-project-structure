# Makefile for the medium-size Go project
-include Makefile.override

BINARY := myapp
CMD := ./cmd/api
BUILD_DIR := bin

MIGRATIONS_DIR := ./migrations
DB_DRIVER ?= postgres

DB_STRING ?= $(DB_URL)
GOOSE := goose

.PHONY: all build run fmt vet test mod-tidy lint clean install \
	migrate-create migrate-up migrate-down migrate-status migrate-redo migrate-rollback migrate-to

all: build

build:
	@mkdir -p $(BUILD_DIR)
	@echo "Building $(BINARY) from $(CMD)..."
	go build -v -o $(BUILD_DIR)/$(BINARY) $(CMD)

run: build
	@echo "Running $(BUILD_DIR)/$(BINARY)..."
	./$(BUILD_DIR)/$(BINARY)

install:
	go install $(CMD)

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

mod-tidy:
	cd myapp && go mod tidy

lint:
	@which golangci-lint > /dev/null 2>&1 || { echo "golangci-lint not found (install it to run lint)"; exit 0; }
	golangci-lint run ./...

# Migrations (goose)
# Usage:
#  - create: make migrate-create NAME=add_users_table
#  - up:     make migrate-up DB_URL="postgres://user:pass@localhost:5432/db?sslmode=disable"
#  - down:   make migrate-down DB_URL=...
#  - status: make migrate-status DB_URL=...
#  - redo:   make migrate-redo DB_URL=...
#  - rollback: make migrate-rollback STEPS=1 DB_URL=...
#  - to version: make migrate-to VERSION=202501011200 DB_URL=...

migrate-create:
	@if [ -z "$(NAME)" ]; then echo "Usage: make migrate-create NAME=your_migration_name"; exit 1; fi
	@$(GOOSE) -dir $(MIGRATIONS_DIR) create $(NAME) sql

migrate-up:
	@echo "Running migrations (up)..."
	@$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_STRING)" up

migrate-down:
	@echo "Running migrations (down)..."
	@$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_STRING)" down

migrate-status:
	@$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_STRING)" status

migrate-redo:
	@$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_STRING)" redo

migrate-rollback:
	@if [ -z "$(STEPS)" ]; then echo "Usage: make migrate-rollback STEPS=1"; exit 1; fi
	@$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_STRING)" rollback $(STEPS)

migrate-to:
	@if [ -z "$(VERSION)" ]; then echo "Usage: make migrate-to VERSION=<version>"; exit 1; fi
	@$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_STRING)" up-to $(VERSION)

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
