#!make

run/gin: 
	go run cmd/server/gin/main.go

run/echo: 
	go run cmd/server/echo/main.go