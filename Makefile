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

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/yuttana76/simbplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock