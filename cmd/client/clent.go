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
	con, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewContactsClient(con)
	res, err := c.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
