up:
	docker compose up -d
build:
	docker compose build --no-cache --force-rm
stop:
	docker compose stop
down:
	docker compose down --remove-orphans
down-v:
	docker compose down --remove-orphans --volumes
destroy:
	docker compose down --rmi all --volumes --remove-orphans
ps:
	docker compose ps
logs:
	docker compose logs
logs-watch:
	docker compose logs --follow
init:
	docker compose up -d --build
	docker compose exec go cp .env.example .env
	@make download-go-module
run-dynamodb:
	docker compose exec go go run src/dynamodb.go
run-lint:
	docker compose exec go golangci-lint run -v
run-fix:
	docker compose exec go golangci-lint run --fix
run-test:
	docker compose exec go go test -cover -v ./...
update-go-module:
	docker compose exec go go mod tidy
download-go-module:
	docker compose exec go go mod download
