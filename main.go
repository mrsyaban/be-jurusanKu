package main

import (
	"JurusanKu/src/api"
	"JurusanKu/src/config"
	db "JurusanKu/src/database/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)


func main() {
	config, err := config.Init(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(*store, config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.SERVER_ADDR)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}