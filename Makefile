# TODO This make file is super slow. Need to optimize it.

help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo 'make dev: make dev for development work'
	@echo 'make build: make build container'
	@echo 'make production: docker production build'
	@echo 'clean: clean for all clear docker images'

init:
	@ if [ -f config.env ]; then echo "config.env already exists!"; exit 1; fi
	@ echo "Copy the config.env.example. to the real config.env file..."
	@ cp config.env.example config.env
	@ echo "Configuration done!"

dev:
	docker-compose -f docker-compose-dev.yml down
	if [ ! -f config.env ]; then cp config.env.example config.env; fi;
	docker-compose --env-file config.env -f docker-compose-dev.yml up

build:
	docker-compose -f docker-compose-prod.yml build
	docker-compose -f docker-compose-dev.yml build

production:
	docker-compose -f docker-compose-prod.yml up -d --build

stop:
	docker-compose -f docker-compose-prod.yml down -v
	docker-compose -f docker-compose-dev.yml down -v

run:
	go run main.go

## I think we can simply use docker to start some of our services, docker compose is abit over kill at this point since it kinda slow and
## I don't want to start the whole cluster just to test a single funciton.
## You can still use it though, but I will just add some docker commandsfor me here.
.EXPORT_ALL_VARIABLES:
### PG
DB_NAME ?= test_pg_go
DB_USER ?= ucsd
DB_PASSWORD ?= 3515
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_LOG_MODE ?= True

stop-pg:
	@echo "stop postgres..."
	@docker stop tree-hole-backend-pg | true

# docker exec tree-hole-backend-pg pg_isready -U postgres: returns 0 (false) if the container is not running
restart-pg: stop-pg
	@echo "start postgres..."
	@docker run -d --rm --name tree-hole-backend-pg \
				-p $(DB_PORT):5432 \
				-e POSTGRES_DB=$(DB_NAME) \
				-e POSTGRES_USER=$(DB_USER) \
				-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
				postgres:13.4-alpine
	@until (docker exec tree-hole-backend-pg pg_isready -U postgres); do \
		echo "[`date`] Waiting for Postgres to be ready"; \
		sleep 1; \
	done
