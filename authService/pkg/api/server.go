package server

import (
	config "authservice/pkg/conifg"
	"authservice/pkg/pb/admin"
	"authservice/pkg/pb/user"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, adminServer admin.AdminServer, userServer user.UserServer) (*Server, error) {
	list, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	newServer := grpc.NewServer()
	admin.RegisterAdminServer(newServer, adminServer)
	user.RegisterUserServer(newServer, userServer)
	return &Server{
		server:   newServer,
		listener: list,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50052")
	return c.server.Serve(c.listener)
}
