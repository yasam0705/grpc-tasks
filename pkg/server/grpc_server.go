package server

import (
	"context"
	"database/sql"
	"fmt"
	"grpc_postgres/pkg/api/proto"

	_ "github.com/lib/pq"
)

const connStr = "user=postgres dbname=psql_task sslmode=disable"

var Db, _ = sql.Open("postgres", connStr)

type GRPCServer struct {
	proto.UnimplementedContactsServer
}

/*
Create(context.Context, *ContactRequest) (*ContactResponse, error)
Update(context.Context, *ContactRequest) (*ContactResponse, error)
Delete(context.Context, *ContactIdRequest) (*ContactResponse, error)
Get(context.Context, *ContactIdRequest) (*ContactResponse, error)
GetAll(context.Context, *ContactIdRequest) (*ContactResponse, error)
*/

func (grpc GRPCServer) Create(ctx context.Context, cr *proto.ContactRequest) (*proto.ContactResponse, error) {
	_, err := Db.Exec(`INSERT INTO contacts
	(contact_id, first_name, last_name, phone, email)
	VALUES ($1, $2, $3, $4, $5)
	`, cr.C.Id, cr.C.FirstName, cr.C.LastName, cr.C.Phone, cr.C.Email)

	return &proto.ContactResponse{C: cr.C}, err
}

func (grpc GRPCServer) Update(ctx context.Context, cr *proto.ContactRequest) (*proto.ContactResponse, error) {
	res, err := Db.Exec(`UPDATE contacts 
			SET first_name = $1, last_name = $2, phone = $3, email = $4
			WHERE contact_id = $5
			`, cr.C.FirstName, cr.C.LastName, cr.C.Phone, cr.C.Email, cr.C.Id)

	if err != nil {
		return &proto.ContactResponse{}, err
	}

	if num, _ := res.RowsAffected(); num == 0 {
		return &proto.ContactResponse{}, fmt.Errorf("contact %d not exists", cr.C.Id)
	}

	return &proto.ContactResponse{C: cr.C}, err
}

func (grpc GRPCServer) Delete(ctx context.Context, cir *proto.ContactIdRequest) (*proto.ContactResponse, error) {
	Db.QueryRow("DELETE FROM contacts WHERE contact_id = $1", cir.Id)

	return &proto.ContactResponse{}, nil
}

func (grpc GRPCServer) Get(ctx context.Context, cir *proto.ContactIdRequest) (*proto.ContactResponse, error) {
	row := Db.QueryRow("SELECT contact_id, first_name, last_name, phone, email from contacts WHERE contact_id = $1", cir.Id)

	var cont proto.Contact
	row.Scan(&cont.Id, &cont.FirstName, &cont.LastName, &cont.Phone, &cont.Email)

	return &proto.ContactResponse{C: &cont}, nil
}

func (grpc GRPCServer) GetAll(ctx context.Context, cir *proto.ContactIdRequest) (*proto.ContactResponse, error) {
	return &proto.ContactResponse{}, nil
}
