DB_NAME = simple_bank
DB_USER = postgres
DB_HOST = localhost
DB_PORT = 5432
DB_PASSWORD = admin

DB_URL = postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

createdb:
	docker exec -it postgres createdb --username=$(DB_USER) --owner=$(DB_USER) simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database $(DB_URL) -verbose up

migratedown:
	migrate -path db/migration -database $(DB_URL) -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown sqlc test