package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/yuttana76/simbplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

// #2
func TestMain(m *testing.M) {
	// var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	//Print connection
	// fmt.Println("Connecting to database...")

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}

// // #1
// func TestMain(m *testing.M) {

// 	conn, err := sql.Open(dbDriver, dbSource)

// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}
// 	testQueries = New(conn)

// 	os.Exit(m.Run())
// }
