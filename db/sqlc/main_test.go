package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yuttana76/simbplebank/util"
)

var testStore Store

// #2
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testStore = NewStore(connPool)

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
