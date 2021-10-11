package main

import (
	"grpc_postgres/pkg/proto"
	"grpc_postgres/pkg/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cs := grpc.NewServer()
	srvc := server.GRPCContactsServer{}

	proto.RegisterContactsServer(cs, srvc)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = cs.Serve(l); err != nil {
		log.Fatal(err)
	}

	ts := grpc.NewServer()
	srvt := server.GRPCTasksServer{}

	proto.RegisterTasksServer(ts, srvt)

	lt, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}

	if err = ts.Serve(lt); err != nil {
		log.Fatal(err)
	}

}
