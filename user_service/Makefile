include app.env

build:
	go build \
		-o build/app \
		cmd/app/main.go

postgres:
	docker run --name lifeat-postgres -p 5435:5435 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it lifeat-postgres createdb --username root --owner root lifeat

dropdb:
	docker exec -it lifeat-postgres dropdb lifeat

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrateup:
	migrate -path migrations -database 'postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable' up

migratedown:
	migrate -path migrations -database 'postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable' down

server:
	go run cmd/app/main.go

.PHONY: build postgres createdb dropdb migrate-create migrateup migratedown server 