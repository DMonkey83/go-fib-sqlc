package main

import (
	"database/sql"
	"log"

	"github.com/DMonkey83/go-fib-sqlc-pg/api"
	db "github.com/DMonkey83/go-fib-sqlc-pg/db/sqlc"
	"github.com/DMonkey83/go-fib-sqlc-pg/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connection, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connection)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
}
