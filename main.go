package main

import (
	"database/sql"
	"log"

	"github.com/harissucipto/xendit-task/api"
	db "github.com/harissucipto/xendit-task/db/sqlc"
	"github.com/harissucipto/xendit-task/util"
	_ "github.com/lib/pq"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:")
	}
	
}