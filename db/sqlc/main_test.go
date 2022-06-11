package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/codeninjaug/simple_bank/db/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:8bdc7axyzex@localhost:5432/swaggie_bank?sslmode=disable"
// )

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	//var err error
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot connect to the configuration files", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
