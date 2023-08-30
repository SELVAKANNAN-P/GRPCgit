package interfaces

import "Netxd_Project1/Netxd_Customer_dal/models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
}
