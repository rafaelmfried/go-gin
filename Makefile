# Makefile para facilitar comandos comuns

.PHONY: help build run test clean docker-build docker-up docker-down docker-logs migrate

# Configurações
APP_NAME=gin-quickstart
DOCKER_COMPOSE=docker-compose

help: ## Mostra esta ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Desenvolvimento local
build: ## Build da aplicação
	go build -o tmp/main ./cmd/api

run: ## Executa a aplicação localmente
	go run ./cmd/api

test: ## Executa os testes
	go test -v ./...

test-coverage: ## Executa testes com coverage
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

clean: ## Limpa arquivos temporários
	rm -rf tmp/
	go clean

# Docker
docker-build: ## Build da imagem Docker
	docker build -t $(APP_NAME) .

docker-up: ## Sobe os containers
	$(DOCKER_COMPOSE) up -d

docker-up-dev: ## Sobe os containers com pgAdmin (modo dev)
	$(DOCKER_COMPOSE) --profile dev up -d

docker-down: ## Para os containers
	$(DOCKER_COMPOSE) down

docker-down-volumes: ## Para os containers e remove volumes
	$(DOCKER_COMPOSE) down -v

docker-logs: ## Mostra logs dos containers
	$(DOCKER_COMPOSE) logs -f

docker-logs-api: ## Mostra logs apenas da API
	$(DOCKER_COMPOSE) logs -f api

docker-restart: ## Reinicia os containers
	$(DOCKER_COMPOSE) restart

docker-rebuild: ## Rebuilda e reinicia os containers
	$(DOCKER_COMPOSE) up -d --build

# Banco de dados
migrate: ## Executa migrações do banco
	go run ./cmd/migrate

seed: ## Executa seeds do banco
	go run ./cmd/seed

# Utilitários
deps: ## Atualiza dependências
	go mod tidy
	go mod download

lint: ## Executa linter
	golangci-lint run

format: ## Formata código
	go fmt ./...
	goimports -w .

# Setup inicial
setup: ## Setup inicial do projeto
	cp .env.example .env
	go mod tidy
	@echo "Projeto configurado! Execute 'make docker-up' para iniciar."