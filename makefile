#!make

run/gin: 
	go run cmd/server/gin/main.go

docker/up: ## Run docker compose
	docker-compose up -d

docker/down: ## Stop docker compose
	docker-compose down
