CONFIG_FILE ?= ./config/local.yml
DB_STRING ?= $(shell sed -n 's/db:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
APP_DSN = $(shell echo $(DB_STRING) | sed -n 's/^[ \t]*//')
MIGRATE := docker run --rm -v $(shell pwd)/migrations:/migrations --network host migrate/migrate:v4.15.1 -path=/migrations/ -database "$(APP_DSN)"

run:
	go run ./cmd/main.go -config ./config/local.yml

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir /migrations/ $${name// /_}

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up


