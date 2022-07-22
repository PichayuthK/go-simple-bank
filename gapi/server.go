package gapi

import (
	"fmt"

	db "github.com/PichayuthK/go-simple-bank/db/sqlc"
	"github.com/PichayuthK/go-simple-bank/pb"
	"github.com/PichayuthK/go-simple-bank/token"
	"github.com/PichayuthK/go-simple-bank/util"
)

// Server servers gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
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
