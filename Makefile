.PHONY: build

-include build/.env

COMPOSE_FILE:=./build/compose.yml

CMD_ARGS?=$(filter-out $@, $(MAKECMDGOALS)) $(MAKEFLAGS)
%:
	@true

up:
	docker compose -f $(COMPOSE_FILE) up -d

down:
	docker compose -f $(COMPOSE_FILE) down

restart: down up

build:
	docker compose -f $(COMPOSE_FILE) build

rebuild:
	docker compose -f $(COMPOSE_FILE) up -d --no-deps --build $(CMD_ARGS)

logs:
	docker compose -f $(COMPOSE_FILE) logs -f
