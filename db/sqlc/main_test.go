package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDrive  = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDrive, dbSource)

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
