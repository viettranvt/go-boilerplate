#!make

run: 
	go run cmd/server/main.go

docker/up: ## Run docker compose
	docker-compose up -d

docker/down: ## Stop docker compose
	docker-compose down

run/deployment: ##run deployment script
	go run ./cmd/deployment/main.go