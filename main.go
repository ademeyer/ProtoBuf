package main

import (
	"log"
	"os"

	contact "github.com/ademeyer/Protocol-Buffer/contact/ProtoFile"
	faker "github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// Serialization: write serialized data
	p1 := contact.Person{
		Name:  faker.Name(),
		Id:    faker.Int32(),
		Email: faker.Email(),
		Phones: []*contact.Person_PhoneNumber{
			{
				Number: faker.Phone(),
				Type:   contact.PhoneType_PHONE_TYPE_MOBILE,
			},
		},
		LastUpdated: timestamppb.Now(),
	}

	p2 := contact.Person{
		Name:  faker.Name(),
		Id:    faker.Int32(),
		Email: faker.Email(),
		Phones: []*contact.Person_PhoneNumber{
			{
				Number: faker.Phone(),
				Type:   contact.PhoneType_PHONE_TYPE_MOBILE,
			},
		},
		LastUpdated: timestamppb.Now(),
	}

	addressbook := &contact.AddressBook{
		People: []*contact.Person{
			&p1,
			&p2,
		},
	}

	log.Println(addressbook)

	out, err := proto.Marshal(addressbook)
	if err != nil {
		log.Fatalln(err)
	}
	_ = os.WriteFile("address.db", out, 0644)

	// Deserialization: read serialized data
	data, err := os.ReadFile("address.db")

	if err != nil {
		log.Fatalln(err)
	}
	var address_book contact.AddressBook
	//address_book := contact.AddressBook{}
	err = proto.Unmarshal(data, &address_book)
	if err != nil {
		log.Fatalln(err)
	}

	for i, person := range address_book.People {
		log.Printf("%d -> %v", i, person)
	}

}
