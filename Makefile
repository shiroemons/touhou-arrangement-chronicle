.PHONY: help init build build-local db-up up down logs ps test migrate generate all-clean
.DEFAULT_GOAL := help

DOCKER_TAG := latest

init: ## Initialize environment
	docker compose build

build: ## Build docker image to deploy
	docker build -t shiroemons/touhou-arrangement-chronicle:${DOCKER_TAG} \
 		--target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

db-up: ## Run docker compose up db
	docker compose up db -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	go test -race -shuffle=on ./...

migrate: ## db migrate
	docker compose run --rm migrate

generate: ## Run go generate ./...
	docker compose run server go generate ./...

all-clean:
	docker compose down --rmi all --volumes --remove-orphans

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
