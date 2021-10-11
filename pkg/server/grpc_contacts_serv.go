package server

import (
	"context"
	"database/sql"
	"fmt"
	"grpc_postgres/pkg/proto"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/emptypb"
)

const connStr = "user=postgres dbname=psql_task sslmode=disable"

var Db, _ = sql.Open("postgres", connStr)

type GRPCContactsServer struct {
	proto.UnimplementedContactsServer
}

func (grpc GRPCContactsServer) Create(ctx context.Context, cr *proto.ContactRequest) (*proto.ContactResponse, error) {
	_, err := Db.Exec(`INSERT INTO contacts
	(contact_id, first_name, last_name, phone, email)
	VALUES ($1, $2, $3, $4, $5)
	`, cr.C.Id, cr.C.FirstName, cr.C.LastName, cr.C.Phone, cr.C.Email)

	return &proto.ContactResponse{C: cr.C}, err
}

func (grpc GRPCContactsServer) Update(ctx context.Context, cr *proto.ContactRequest) (*proto.ContactResponse, error) {
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

func (grpc GRPCContactsServer) Delete(ctx context.Context, cir *proto.ContactIdRequest) (*proto.ContactResponse, error) {
	Db.QueryRow("DELETE FROM contacts WHERE contact_id = $1", cir.Id)

	return &proto.ContactResponse{}, nil
}

func (grpc GRPCContactsServer) Get(ctx context.Context, cir *proto.ContactIdRequest) (*proto.ContactResponse, error) {
	row := Db.QueryRow("SELECT contact_id, first_name, last_name, phone, email from contacts WHERE contact_id = $1", cir.Id)

	var cont proto.Contact
	row.Scan(&cont.Id, &cont.FirstName, &cont.LastName, &cont.Phone, &cont.Email)

	return &proto.ContactResponse{C: &cont}, nil
}

func (grpc GRPCContactsServer) GetAll(ctx context.Context, emp *emptypb.Empty) (*proto.ContactSliceResponce, error) {
	rows, err := Db.Query("SELECT contact_id, first_name, last_name, phone, email from contacts ORDER BY contact_id")
	if err != nil {
		return &proto.ContactSliceResponce{C: make([]*proto.Contact, 0)}, err
	}

	defer rows.Close()

	var result = make([]*proto.Contact, 0)

	for rows.Next() {
		var tempContact proto.Contact
		rows.Scan(&tempContact.Id, &tempContact.FirstName, &tempContact.LastName, &tempContact.Phone, &tempContact.Email)

		result = append(result, &tempContact)
	}

	return &proto.ContactSliceResponce{C: result}, nil
}
