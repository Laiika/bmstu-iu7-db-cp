SWAGGER_SRC := $(wildcard ./internal/server/*.go)

swagger: $(SWAGGER_SRC)
	swag init --parseDependency --parseInternal -g ./cmd/app/main.go -o ./swagger

build:
	go build ./cmd/app/main.go
	./main.exe

integrational-tests:
	go test ./tests/integrational -c -o tests.exe
	./tests.exe

unit-tests:
	go test ./internal/domain/service -cover
	go test ./internal/domain/service/auth -cover

run:
	docker-compose up -d

down:
	docker-compose down