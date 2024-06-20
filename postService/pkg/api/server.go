package server

import (
	"fmt"
	"net"
	"postservice/pkg/config"
	"postservice/pkg/pb/post"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, postServer post.PostServer) (*Server, error) {
	fmt.Println(cfg.Port)
	list, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	newServer := grpc.NewServer()
	post.RegisterPostServer(newServer, postServer)
	return &Server{
		server:   newServer,
		listener: list,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50053")
	return c.server.Serve(c.listener)
}
