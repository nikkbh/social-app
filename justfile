# https://github.com/casey/just/blob/master/README.md

set dotenv-load
set positional-arguments

migrate name:
    @migrate create -dir ./cmd/migrate/migrations -seq -ext sql {{name}}

migrate-up:
    @migrate --path ./cmd/migrate/migrations --database $DB_ADDR up

migrate-down:
    @migrate --path ./cmd/migrate/migrations --database $DB_ADDR down

docker-up:
    @ "/C/Program Files/Docker/Docker/Docker Desktop.exe" &

db-up:
    @docker start postgres-db

db-down:
    @docker stop postgres-db

seed:
    @go run ./cmd/migrate/seed/main.go
