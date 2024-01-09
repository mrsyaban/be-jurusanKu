package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	db "JurusanKu/src/database/sqlc"
	"JurusanKu/src/api"
)
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5555/jurusanKu?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(*store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}