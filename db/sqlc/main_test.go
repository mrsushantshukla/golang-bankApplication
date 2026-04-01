package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DBDriverName = "postgres"
	DBConnString = "postgresql://root:root@localhost:5432/venus_bank?sslmode=disable"
)

var testQueries *Queries
var testStore *Store

func TestMain(m *testing.M) {
	conn, err := sql.Open(DBDriverName, DBConnString)

	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testQueries = New(conn)
	testStore = NewStore(conn)

	os.Exit(m.Run())

}
