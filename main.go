package main

import (
	"database/sql"
	"github.com/bjclayton/simplebank/api"
	db "github.com/bjclayton/simplebank/db/sqlc"
	"github.com/bjclayton/simplebank/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDb, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	store := db.NewStore(testDb)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
