package gapi

import (
	"fmt"

	db "github.com/ashiqYousuf/sbank/db/sqlc"
	"github.com/ashiqYousuf/sbank/pb"
	"github.com/ashiqYousuf/sbank/token"
	"github.com/ashiqYousuf/sbank/util"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, nil
}
