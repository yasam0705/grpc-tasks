package tt

import (
	"context"
	"fmt"
	"grpc_postgres/pkg/proto"
	"grpc_postgres/pkg/server"
	"log"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	testContacts = []proto.Contact{
		{
			Id:        1000,
			FirstName: "Sam",
			LastName:  "Smith",
			Phone:     "(695)-175-4661",
			Email:     "sam@local.com",
		},
		{
			Id:        1100,
			FirstName: "Eugene",
			LastName:  "Williamson",
			Phone:     "(139)-191-0039",
			Email:     "eugene@local.com",
		},
		{
			Id:        1200,
			FirstName: "Brian",
			LastName:  "Robinson",
			Phone:     "(045)-207-9455",
			Email:     "brian.robinson@example.com",
		},
	}
	updContact = proto.Contact{
		Id:        1200,
		FirstName: "Lee",
		LastName:  "Wright",
		Phone:     "(215)-511-9272",
		Email:     "lee.wright@example.com",
	}
	clt proto.ContactsClient
)

func NewClient() proto.ContactsClient {
	con, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return proto.NewContactsClient(con)
}

func TestCreateContact(t *testing.T) {
	clt = NewClient()
	var del string

	for i, v := range testContacts {
		res, err := clt.Create(context.Background(), &proto.ContactRequest{C: &v})
		if err != nil {
			t.Error(err)
		}
		if res.GetC().Phone != v.Phone {
			t.Error("failed create method")
		}

		if i == len(testContacts)-1 {
			del += fmt.Sprintf("%d", v.Id)
		} else {
			del += fmt.Sprintf("%d, ", v.Id)
		}
	}

	t.Cleanup(func() {
		delQuery := fmt.Sprintf("DELETE FROM contacts WHERE contact_id IN (%s)", del)
		_, err := server.Db.Exec(delQuery)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestUpdateContact(t *testing.T) {
	TestCreateContact(t)

	res, err := clt.Update(context.Background(), &proto.ContactRequest{C: &updContact})
	if err != nil {
		t.Error(err)
	}
	if res.GetC().Phone != updContact.Phone {
		t.Error("failed update method")
	}
}

func TestGetContact(t *testing.T) {
	TestCreateContact(t)

	for _, v := range testContacts {
		temp, err := clt.Get(context.Background(), &proto.ContactIdRequest{Id: v.GetId()})
		if err != nil {
			t.Error(err)
		}
		if temp.GetC().Id != v.Id {
			t.Error("method get failed")
		}
	}
}

func TestGetAllContact(t *testing.T) {
	TestCreateContact(t)

	temp, err := clt.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Error(err)
	}
	for i := range testContacts {
		if testContacts[i].Id != temp.C[i].Id {
			t.Error("failed getall method")
		}
	}
}

func TestGetDeleteContact(t *testing.T) {
	TestCreateContact(t)

	for _, v := range testContacts {
		_, err := clt.Delete(context.Background(), &proto.ContactIdRequest{Id: v.GetId()})
		if err != nil {
			t.Error(err)
		}
	}
}
