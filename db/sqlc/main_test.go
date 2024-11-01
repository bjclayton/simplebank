package db

import (
	"database/sql"
	"github.com/bjclayton/simplebank/util"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	testQueries = New(testDb)

	code := m.Run()

	if err := testDb.Close(); err != nil {
		log.Fatal("Failed to close the database connection: ", err)
	}
	os.Exit(code)
}
