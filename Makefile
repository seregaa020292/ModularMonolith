.PHONY: build init

-include ./build/.env

COMPOSE_FILE:=./build/compose.yml
MIGRATION_DIR:=./migrations
LOCAL_BIN:=$(CURDIR)/bin

CMD_ARGS?=$(filter-out $@, $(MAKECMDGOALS)) $(MAKEFLAGS)
%:
	@true

up:
	docker compose -f $(COMPOSE_FILE) up -d

down:
	docker compose -f $(COMPOSE_FILE) down

restart: down up

init: env-create install-deps build migration-up

build:
	docker compose -f $(COMPOSE_FILE) build

rebuild:
	docker compose -f $(COMPOSE_FILE) up -d --no-deps --build $(CMD_ARGS)

logs:
	docker compose -f $(COMPOSE_FILE) logs -f

lint:
	$(LOCAL_BIN)/golangci-lint run

test:
	docker build -f ./build/Dockerfile --target test .

test-export:
	docker build -f ./build/Dockerfile --target test-export -q -o ./tmp/test .

migration-create:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) create $(CMD_ARGS) sql

migration-status:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) postgres ${PG_DSN} status -v

migration-up:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) postgres ${PG_DSN} up -v

migration-down:
	$(LOCAL_BIN)/goose -dir $(MIGRATION_DIR) postgres ${PG_DSN} down -v

generate-openapi:
	mkdir -p ./internal/ports/oapi
	$(LOCAL_BIN)/oapi-codegen -generate chi-server -package oapi -o ./internal/ports/oapi/openapi_server.go ./pkg/specs/openapi/swagger.yml
	$(LOCAL_BIN)/oapi-codegen -generate types -package oapi -o ./internal/ports/oapi/openapi_types.go ./pkg/specs/openapi/swagger.yml
	#$(LOCAL_BIN)/oapi-codegen -generate client -package oapi -o ./internal/ports/oapi/openapi_client.go ./pkg/specs/openapi/swagger.yml

generate-jet:
	$(LOCAL_BIN)/jet -source=postgres -dsn=${PG_DSN} -path=./internal/models -ignore-tables=goose_db_version

generate-wire:
	$(LOCAL_BIN)/wire ./internal/infrastructure/app

env-create:
	[ -f ./build/.env ] || cp ./build/.env.example ./build/.env

install-deps:
	[ -f $(LOCAL_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN) v1.59.0
	[ -f $(LOCAL_BIN)/goose ] || curl -sSfL https://raw.githubusercontent.com/pressly/goose/master/install.sh | GOOSE_INSTALL=. sh -s v3.20.0
	[ -f $(LOCAL_BIN)/oapi-codegen ] || GOBIN=$(LOCAL_BIN) go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0
	[ -f $(LOCAL_BIN)/jet ] || GOBIN=$(LOCAL_BIN) go install github.com/go-jet/jet/v2/cmd/jet@v2.11.1
	[ -f $(LOCAL_BIN)/wire ] || GOBIN=$(LOCAL_BIN) go install github.com/google/wire/cmd/wire@v0.6.0
