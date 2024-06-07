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
	docker compose -f $(COMPOSE_FILE) logs -f $(CMD_ARGS)

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

generate-oapi-server:
	OUTPUT_DIR=./internal/infrastructure/openapi ; \
 	SPEC_FILE=./pkg/specs/openapi/swagger.yml ; \
	$(LOCAL_BIN)/oapi-codegen -generate chi-server,strict-server -package openapi -o $$OUTPUT_DIR/openapi_server.go $$SPEC_FILE ; \
	$(LOCAL_BIN)/oapi-codegen -generate types -package openapi -o $$OUTPUT_DIR/openapi_types.go $$SPEC_FILE ; \
	$(LOCAL_BIN)/oapi-codegen -generate spec -package openapi -o $$OUTPUT_DIR/openapi_spec.go $$SPEC_FILE

generate-oapi-client:
	#$(LOCAL_BIN)/oapi-codegen -generate client -package openapi -o $$OUTPUT_DIR/openapi_client.go $$SPEC_FILE

generate-jet:
	$(LOCAL_BIN)/jet -source=postgres -dsn=${PG_DSN} -path=./internal/models -ignore-tables=goose_db_version

generate-wire:
	$(LOCAL_BIN)/wire ./internal/infrastructure/app

env-create:
	[ -f ./build/.env ] || cp ./build/.env.example ./build/.env

install-deps:
	[ -f $(LOCAL_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN) v1.59.0
	[ -f $(LOCAL_BIN)/goose ] || curl -sSfL https://raw.githubusercontent.com/pressly/goose/master/install.sh | GOOSE_INSTALL=. sh -s v3.20.0
	[ -f $(LOCAL_BIN)/oapi-codegen ] || GOBIN=$(LOCAL_BIN) go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0
	[ -f $(LOCAL_BIN)/jet ] || GOBIN=$(LOCAL_BIN) go install github.com/go-jet/jet/v2/cmd/jet@v2.11.1
	[ -f $(LOCAL_BIN)/wire ] || GOBIN=$(LOCAL_BIN) go install github.com/google/wire/cmd/wire@v0.6.0
