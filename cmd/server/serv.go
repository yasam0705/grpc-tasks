package main

import (
	"grpc_postgres/pkg/proto"
	"grpc_postgres/pkg/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	go contactServ()
	taskServ()

}
func taskServ() {
	ts := grpc.NewServer()
	srvt := server.GRPCTasksServer{}

	proto.RegisterTasksServer(ts, srvt)

	lt, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	if err = ts.Serve(lt); err != nil {
		log.Fatal(err)
	}
}

func contactServ() {
	cs := grpc.NewServer()
	srvc := server.GRPCContactsServer{}

	proto.RegisterContactsServer(cs, srvc)

	lc, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = cs.Serve(lc); err != nil {
		log.Fatal(err)
	}
}
