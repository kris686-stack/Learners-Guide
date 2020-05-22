package main

import (
	"io"
	"log"

	"context"
	"google.golang.org/grpc"

	"agrpc"
)

const (
	address = "localhost:5051"
)

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createCustomer(client agrpc.CustomerClient, customer *agrpc.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getCustomers(client agrpc.CustomerClient, filter *agrpc.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", customer)
	}
}
func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := agrpc.NewCustomerClient(conn)

	customer := &agrpc.CustomerRequest{
		Id:    101,
		Name:  "Krishna Raj",
		Email: "krishnaraj686@gmail.com",
		Phone: "9526378279",
		Addresses: []*agrpc.CustomerRequest_Address{
			&agrpc.CustomerRequest_Address{
				Street:            "Kannan Villai",
				City:              "Malayalapuzha ,Pathanamthitta",
				State:             "Kerala",
				Zip:               "689666",
				IsShippingAddress: false,
			},
			&agrpc.CustomerRequest_Address{
				Street:            "7g Bhoomi Reddy Colony",
				City:              "New Tippassandra ,Banglore",
				State:             "Karnataka",
				Zip:               "650075",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)

	customer=&agrpc.CustomerRequest{

		Id:    102,
		Name:  "Unni",
		Email: "ubalu@live.com",
		Phone: "56839",
		Addresses: []*agrpc.CustomerRequest_Address{
			&agrpc.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "missisuga",
				State:             "Onatrio",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)
	// Filter with an empty Keyword
	filter :=&agrpc.CustomerFilter{Keyword:""}
	getCustomers(client, filter)
}