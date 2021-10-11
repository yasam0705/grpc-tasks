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
	// contactServer, err := grpc.Dial(":8080", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cs := proto.NewContactsClient(contactServer)
	// res, err := cs.GetAll(context.Background(), &emptypb.Empty{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

	taskServer, err := grpc.Dial(":8181", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	ts := proto.NewTasksClient(taskServer)
	ts.Create(context.Background(), &proto.TaskRequest{T: &proto.Task{
		Id:        10,
		Name:      "Task 10",
		Status:    "done",
		Priority:  "Important",
		CreatedAt: "04.10.2021",
		CreatedBy: "05.10.2021",
		DueDate:   "08.10.2021",
	}})
	fmt.Println(ts.GetAll(context.Background(), &emptypb.Empty{}))
}
