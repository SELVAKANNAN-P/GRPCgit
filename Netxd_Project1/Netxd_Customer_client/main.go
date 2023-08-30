package main

import (
	"context"
	"fmt"
	"log"

	pb "Netxd_Project1/customer"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect :%v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pb.Customer{
		CustomerID: 123,
		FirstName:  "kishore",
		LastName:   "N S",
		BankID:     123,

		UpdatedAt: "2023.08.30",
		IsActive:  true,
	})
	if err != nil {
		log.Fatalf("Failed to create customer :%v", err)
	}

	fmt.Printf("Response from customer : %v\n", response.CustomerID)
}
