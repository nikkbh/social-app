# https://github.com/casey/just/blob/master/README.md

set dotenv-load
set positional-arguments

migrate name:
    @migrate create -dir ./cmd/migrate/migrations -seq -ext sql {{name}}

migrate-up:
    @migrate --path ./cmd/migrate/migrations --database $DB_ADDR up

migrate-down:
    @migrate --path ./cmd/migrate/migrations --database $DB_ADDR down

migrate-force version:
    @migrate --path ./cmd/migrate/migrations --database $DB_ADDR force {{version}}

docker:
    @ "/C/Program Files/Docker/Docker/Docker Desktop.exe" &

db-up:
    @docker start postgres-db

db-down:
    @docker stop postgres-db

redis-up:
    @docker start social-redis

redis-down:
    @docker stop social-redis

storage-up:
    @docker start social-redis
    @docker start postgres-db

storage-down:
    @docker stop social-redis 
    @docker stop postgres-db

seed:
    @go run ./cmd/migrate/seed/main.go

gen-docs:
    @swag init -g ./api/main.go -d cmd,internal && swag fmt

test:
    @go test -v ./...