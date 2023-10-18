.PHONY: build up down logs

build:
	docker compose -f ./deploy/docker-compose.yml up -d --remove-orphans

up:
	docker compose -f ./deploy/docker-compose.yml up -d

down:
	docker compose -f ./deploy/docker-compose.yml down

logs:
	docker compose -f ./deploy/docker-compose.yml logs