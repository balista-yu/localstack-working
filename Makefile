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
run-main:
	docker compose exec go go run src/main.go
run-lint:
	docker compose exec go golangci-lint run -v
run-fix:
	docker compose exec go golangci-lint run --fix
