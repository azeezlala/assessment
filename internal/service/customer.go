package service

import (
	"context"
	"errors"
	"github.com/azeezlala/assessment/internal/model"
	"github.com/azeezlala/assessment/internal/repository"
	"github.com/google/uuid"
	"log"
)

type ICustomerService interface {
	CreateCustomer(ctx context.Context, data model.Customer) (*model.Customer, error)
}

type CustomerService struct {
	customer repository.ICustomerRepository
}

func NewCustomerService(customer repository.ICustomerRepository) ICustomerService {
	return &CustomerService{customer: customer}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, data model.Customer) (*model.Customer, error) {
	log.Printf("creating customer with email: %v", data.Email)

	if len(data.Email) < 1 {
		return nil, errors.New("email is required")
	}

	if len(data.Name) < 1 {
		return nil, errors.New("name is required")
	}

	data.ID = uuid.NewString()
	res, err := s.customer.CreateCustomer(ctx, &data)
	if err != nil {
		log.Printf("unable to create customer: %v", err)
		return nil, errors.New("unable to create customer")
	}

	return res, nil
}
