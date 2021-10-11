package main

import (
	"context"
	"fmt"
	"grpc_postgres/pkg/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	contactServer, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	cs := proto.NewContactsClient(contactServer)
	res, err := cs.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	taskServer, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	ts := proto.NewTasksClient(taskServer)
	rest, err := ts.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rest)
}
