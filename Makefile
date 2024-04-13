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
	@ echo "Copy the .env.example. to the real .env file..."
	@ cp .env.example .env
	@ echo "Configuration done!"

dev:
	docker-compose -f docker-compose-dev.yml down
	if [ ! -f .env ]; then cp .env.example .env; fi;
	docker-compose -f docker-compose-dev.yml up

build:
	docker-compose -f docker-compose-prod.yml build
	docker-compose -f docker-compose-dev.yml build

production:
	docker-compose -f docker-compose-prod.yml up -d --build

stop:
	docker-compose -f docker-compose-prod.yml down -v
	docker-compose -f docker-compose-dev.yml down -v

