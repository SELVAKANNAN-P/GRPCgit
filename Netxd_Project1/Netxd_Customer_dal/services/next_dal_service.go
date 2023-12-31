package services

import (
	"Netxd_Project1/Netxd_Customer_dal/interfaces"
	"Netxd_Project1/Netxd_Customer_dal/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (p *CustomerService) CreateCustomer(customer *models.Customer) (*models.DBResponse, error) {
	// Insert the customer into the MongoDB collection
	_, err := p.CustomerCollection.InsertOne(p.ctx, customer)
	if err != nil {
		return nil, err
	}

	// Construct the response with customer ID and CreatedAt
	response := &models.DBResponse{
		CustomerID: customer.CustomerID,
		CreatedAt:  customer.CreatedAt,
	}

	return response, nil
}
