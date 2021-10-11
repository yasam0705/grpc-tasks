package server

import (
	"context"
	"fmt"
	"grpc_postgres/pkg/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCTasksServer struct {
	proto.UnimplementedTasksServer
}

/*
Create(context.Context, *TaskRequest) (*TaskResponse, error)
Update(context.Context, *TaskRequest) (*TaskResponse, error)
Delete(context.Context, *TaskIdRequest) (*TaskResponse, error)
Get(context.Context, *TaskIdRequest) (*TaskResponse, error)
GetAll(context.Context, *emptypb.Empty) (*TaskSliceResponce, error)
mustEmbedUnimplementedTasksServer()
*/

func (grpc GRPCTasksServer) Create(ctx context.Context, tr *proto.TaskRequest) (*proto.TaskResponse, error) {
	_, err := Db.Exec(`INSERT INTO tasks
	(task_id, title, status, priority, created_at, created_by, due_date)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, tr.T.GetId(), tr.T.GetName(), tr.T.GetStatus(), tr.T.GetPriority(), tr.T.GetCreatedAt(), tr.T.GetCreatedBy(), tr.T.GetDueDate())

	return &proto.TaskResponse{T: tr.T}, err
}

func (grpc GRPCTasksServer) Update(ctx context.Context, tr *proto.TaskRequest) (*proto.TaskResponse, error) {
	res, err := Db.Exec(`UPDATE tasks 
			SET title = $1, status = $2, priority = $3, due_date = $4
			WHERE task_id = $5
			`, tr.T.GetName(), tr.T.GetStatus(), tr.T.GetPriority(), tr.T.GetDueDate(), tr.T.GetId())

	if err != nil {
		return &proto.TaskResponse{}, err
	}

	if num, _ := res.RowsAffected(); num == 0 {
		return &proto.TaskResponse{}, fmt.Errorf("task %d not exists", tr.T.GetId())
	}
	return &proto.TaskResponse{T: tr.T}, nil
}

func (grpc GRPCTasksServer) Get(ctx context.Context, tir *proto.TaskIdRequest) (*proto.TaskResponse, error) {
	row := Db.QueryRow("SELECT task_id, title, status, priority, created_at, created_by, due_date from tasks WHERE task_id = $1", tir.GetId())

	var task proto.Task
	row.Scan(&task.Id, &task.Name, &task.Status, &task.Priority, &task.CreatedAt, &task.CreatedBy, &task.DueDate)

	return &proto.TaskResponse{T: &task}, nil
}

func (grpc GRPCTasksServer) GetAll(ctx context.Context, tr *emptypb.Empty) (*proto.TaskSliceResponce, error) {
	rows, err := Db.Query("SELECT task_id, title, status, priority, created_at, created_by, due_date from tasks ORDER BY task_id")
	if err != nil {
		return &proto.TaskSliceResponce{T: make([]*proto.Task, 0)}, err
	}

	defer rows.Close()

	var result []*proto.Task

	for rows.Next() {
		var tempTask proto.Task
		rows.Scan(&tempTask.Id, &tempTask.Name, &tempTask.Status, &tempTask.Priority, &tempTask.CreatedAt, &tempTask.CreatedBy, &tempTask.DueDate)

		result = append(result, &tempTask)
	}

	return &proto.TaskSliceResponce{T: result}, nil
}

func (grpc GRPCTasksServer) Delete(ctx context.Context, tir *proto.TaskIdRequest) (*proto.TaskResponse, error) {
	_, err := Db.Exec("DELETE FROM tasks WHERE task_id = $1", tir.GetId())
	return &proto.TaskResponse{}, err
}
