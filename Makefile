ifneq (,$(wildcard ./.env))
    include .env
    export
endif

test:
	go test -v ./...
run:
	go run ./cmd/main.go
create_migrate_%:
	migrate create -ext sql -dir internal/migrations -seq %
migrate:
	migrate -database postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE} -path internal/migrations up