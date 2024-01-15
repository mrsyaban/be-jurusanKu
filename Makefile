db:
	docker run --name db-jurusanku -p 5555:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it db-jurusanku createdb --username=root --owner=root jurusanKu

dropdb:
	docker exec -it db-jurusanku dropdb jurusanKu

migrateup:
	migrate -path src/database/migration -database "postgresql://admin:rakuten@34.128.80.197:5432/db_jurusanku" -verbose up

migratedown:
	migrate -path src/database/migration -database "postgresql://admin:rakuten@34.128.80.197:5432/db_jurusanku" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

opendb:
	docker exec -it db-jurusanku psql -U root jurusanKu

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server