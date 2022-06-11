package main

import (
	"database/sql"
	"log"

	"github.com/codeninjaug/simple_bank/api"
	db "github.com/codeninjaug/simple_bank/db/sqlc"
	"github.com/codeninjaug/simple_bank/db/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:8bdc7axyzex@localhost:5432/swaggie_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil{
		log.Fatal("cannot load the configuration files", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
