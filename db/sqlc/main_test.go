package db

import (
	"database/sql"
	"log"
	"testing"
	"os"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
