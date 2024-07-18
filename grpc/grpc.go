package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RegisterGet(port *int64, server func(server *grpc.Server)) {
	//注册伊特tcp链接 传一个端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
