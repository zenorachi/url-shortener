.SILENT:
.PHONY: proto build run stop rebuild
.DEFAULT_GOAL: run

PATH_TO_PROTO=./api/proto
PROTO_OUT=./pkg/api

proto:
	protoc \
	--go_out=$(PROTO_OUT) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_OUT) \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$(PROTO_OUT) \
	--grpc-gateway_opt=paths=source_relative \
	--proto_path=$(PATH_TO_PROTO) $(PATH_TO_PROTO)/*.proto

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up --remove-orphans

stop:
	docker-compose down

rebuild: build
	docker-compose up --remove-orphans --build

lint:
	golangci-lint run

swag:
	protoc \
	--openapiv2_out ./docs \
	--proto_path=$(PATH_TO_PROTO) $(PATH_TO_PROTO)/*.proto

clean:
	rm -rf .bin/