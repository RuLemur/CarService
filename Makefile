CONFIG_FILE ?= ./config/local.yml
DB_STRING ?= $(shell sed -n 's/db:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
APP_DSN = $(shell echo $(DB_STRING) | xargs)
MIGRATE := docker run --rm -v $(shell pwd)/migrations:/migrations --network host migrate/migrate:v4.15.1 -path=/migrations/ -database "$(APP_DSN)"

run:
	go run ./cmd/main.go -config ./config/local.yml

proto-gen:
	@echo "Running generation proto"
	cp ./internal/app/car_service/endpoint/messages.proto ./pkg/endpoint/messages.proto
	sed -i '' 's/package .*.endpoint;/package endpoint;/' ./pkg/endpoint/messages.proto
	sed -i '' 's|option go_package = ".*endpoint";|option go_package = "pkg/endpoint";|' ./pkg/endpoint/messages.proto
	protoc -I=. -I=$(GOPATH)/src --gofast_out=plugins=grpc:. ./internal/app/car_service/endpoint/messages.proto
	protoc -I=. -I=$(GOPATH)/src --gofast_out=plugins=grpc:. ./pkg/endpoint/messages.proto

start:
	docker-compose up -d
.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@echo $(APP_DSN)
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


