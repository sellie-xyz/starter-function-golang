.PHONY: build
build:
	sam build

.PHONY: run
run: build
	sam local start-api

.PHONY: run-local
run-local:
	PORT=8005 TABLE_NAME=object AWS_REGION=us-east-1 go run cmd/server/main.go
