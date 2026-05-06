package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/yuttana76/simbplebank/api"
	db "github.com/yuttana76/simbplebank/db/sqlc"
	"github.com/yuttana76/simbplebank/gapi"
	"github.com/yuttana76/simbplebank/pb"
	"github.com/yuttana76/simbplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	runGRPCServer(config, store)
	// runGinServer(config, store) //Old for rest api with gin

}

func runGRPCServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)

	// enable reflection to explore functions are aviable in grpc server
	// by using grpcurl
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	log.Printf("start gRPC server: %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}

}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
