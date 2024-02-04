postgres:
	docker run --name simplebankgolang -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=simplebankgolang -e POSTGRES_DB=simplebankgolang -d postgres:16-alpine

createdb:
	docker exec -it simplebankgolang createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it simplebankgolang dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:simplebankgolang@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:simplebankgolang@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test