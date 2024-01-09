postgres:
	docker run --name postgres16 -p 5555:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root jurusanKu

dropdb:
	docker exec -it postgres16 dropdb jurusanKu

migrateup:
	migrate -path src/database/migration -database "postgresql://root:secret@localhost:5555/jurusanKu?sslmode=disable" -verbose up

migratedown:
	migrate -path src/database/migration -database "postgresql://root:secret@localhost:5555/jurusanKu?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server