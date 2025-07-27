APP_NAME=calculator
.PHONY:test

build:
	@echo "building App..."
	@go build -o ${APP_NAME} cmd/main.go
	
run:
	@go run cmd/main.go

test:
	@go test ./...

docker-build:
	@docker build -t calculator-slim .

docker-run:
	@docker run -it calculator-slim


clean:
	@rm -f ${APP_NAME}