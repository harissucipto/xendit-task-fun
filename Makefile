postgres:
	docker run --name postgres12xendit -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12xendit createdb --username=root --owner=root xendit

dropdb:
	docker exec -it postgres12 dropdb xendit

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/xendit?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/xendit?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server