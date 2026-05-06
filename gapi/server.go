package gapi

import (
	"fmt"

	db "github.com/yuttana76/simbplebank/db/sqlc"
	"github.com/yuttana76/simbplebank/pb"
	"github.com/yuttana76/simbplebank/token"
	"github.com/yuttana76/simbplebank/util"
)

// Server serve gRPC request for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	return server, nil
}
