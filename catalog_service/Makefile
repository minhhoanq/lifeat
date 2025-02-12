include app.env

proto:
	buf generate api

migrate-create:
	migrate create -ext sql -dir ./internal/dataaccess/database/migrations/postgres -seq $(name)

migrateup:
	migrate -path ./internal/dataaccess/database/migrations/postgres -database 'postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable' up

migratedown:
	migrate -path ./internal/dataaccess/database/migrations/postgres -database 'postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable' down

server:
	go run cmd/app/main.go

.PHONY: proto migrate-create migrateup migratedown server