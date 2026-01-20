DB_URL=postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker compose up -d db-simplebank

createdb:
	docker exec -it db-simplebank createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it db-simplebank dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown sqlc test server