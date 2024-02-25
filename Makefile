postgres:
	docker run --name simplebankgolang -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=simplebankgolang -e POSTGRES_DB=simplebankgolang -d postgres:16-alpine

createdb:
	docker exec -it simplebankgolang createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it simplebankgolang dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:simplebankgolang@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:simplebankgolang@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:simplebankgolang@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:simplebankgolang@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

mockgen:
	mockgen -package mockdb --build_flags=--mod=mod -destination db/mock/store.go github.com/skylineCodes/bank_app/db/sqlc Store

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc mockgen test server