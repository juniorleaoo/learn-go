.PHONY: default run build test docs

APP_NAME=gopportunities

default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go build =o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init