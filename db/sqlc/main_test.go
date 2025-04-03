package db

import (
	"os"
	"database/sql"
	"log"
	"testing"


	_"github.com/lib/pq"
)

const (
	dbDriver ="postgres"
	dbSource ="postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)

	// Ensure database connection is closed after tests
	defer conn.Close()

	// Run tests only once
	code := m.Run()
	os.Exit(code)
}
