package main

import (
	"grpc_postgres/pkg/proto"
	"grpc_postgres/pkg/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := server.GRPCServer{}

	proto.RegisterContactsServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
