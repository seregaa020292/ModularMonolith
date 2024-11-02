.PHONY: build init

-include ./build/.env

DOCKERFILE_PATH:=./build/Dockerfile
COMPOSE_FILE:=./build/compose.yml
MIGRATION_DIR:=./migrations

CLI_NAME:=cli-$(shell basename "$(CURDIR)")
CLI_DOCKER_EXEC:=docker run --rm --mount type=bind,source=./,target=/app $(CLI_NAME)

CMD_ARGS?=$(filter-out $@, $(MAKECMDGOALS)) $(MAKEFLAGS)
%:
	@true

up:
	docker compose -f $(COMPOSE_FILE) up -d

down:
	docker compose -f $(COMPOSE_FILE) down

restart: down up

init: env-create cli-build build migrate-up

build:
	docker compose -f $(COMPOSE_FILE) build

rebuild:
	docker compose -f $(COMPOSE_FILE) up -d --no-deps --build $(CMD_ARGS)

logs:
	docker compose -f $(COMPOSE_FILE) logs -f $(CMD_ARGS)

cli-build:
	docker build -t $(CLI_NAME) -f $(DOCKERFILE_PATH) --target cli .

cli-export-golangci-lint:
	$(CLI_DOCKER_EXEC) sh -c 'cp $$GOPATH/bin/golangci-lint ./bin/golangci-lint'

lint:
	$(CLI_DOCKER_EXEC) golangci-lint run

test:
	docker build -f $(DOCKERFILE_PATH) --target test .

test-export:
	docker build -f $(DOCKERFILE_PATH) --target test-export -q -o ./tmp/test .

migrate-create:
	$(CLI_DOCKER_EXEC) goose -dir $(MIGRATION_DIR) create $(CMD_ARGS) sql

migrate-status:
	$(CLI_DOCKER_EXEC) goose -dir $(MIGRATION_DIR) postgres ${PG_DSN} status -v

migrate-up:
	$(CLI_DOCKER_EXEC) goose -dir $(MIGRATION_DIR) postgres ${PG_DSN} up -v

migrate-down:
	$(CLI_DOCKER_EXEC) goose -dir $(MIGRATION_DIR) postgres ${PG_DSN} down -v

gen-oapi-server:
	OUTPUT_DIR=./internal/infrastructure/server/openapi ; \
 	SPEC_FILE=./pkg/specification/openapi/swagger.yml ; \
	$(CLI_DOCKER_EXEC) oapi-codegen -generate chi-server,strict-server -package openapi -o $$OUTPUT_DIR/openapi_server.go $$SPEC_FILE ; \
	$(CLI_DOCKER_EXEC) oapi-codegen -generate types -package openapi -o $$OUTPUT_DIR/openapi_types.go $$SPEC_FILE ; \
	$(CLI_DOCKER_EXEC) oapi-codegen -generate spec -package openapi -o $$OUTPUT_DIR/openapi_spec.go $$SPEC_FILE

gen-oapi-client:
	#$(CLI_DOCKER_EXEC) oapi-codegen -generate client -package openapi -o $$OUTPUT_DIR/openapi_client.go $$SPEC_FILE

gen-jet:
	$(CLI_DOCKER_EXEC) jet -source=postgres -dsn=${PG_DSN} -path=./internal/models -ignore-tables=goose_db_version

gen-wire:
	$(CLI_DOCKER_EXEC) wire ./internal/config/di

env-create:
	[ -f ./build/.env ] || cp ./build/.env.example ./build/.env
