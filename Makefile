db:
	docker run --name db-jurusanku -p 5555:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it db-jurusanku createdb --username=root --owner=root jurusanKu

dropdb:
	docker exec -it db-jurusanku dropdb jurusanKu

migrateup:
	migrate -path src/database/migration -database "postgresql://root:secret@localhost:5555/jurusanKu?sslmode=disable" -verbose up

migratedown:
	migrate -path src/database/migration -database "postgresql://root:secret@localhost:5555/jurusanKu?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

opendb:
	docker exec -it db-jurusanku psql -U root jurusanKu

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server