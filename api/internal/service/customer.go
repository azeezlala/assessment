package service

import (
	"context"
	"errors"
	"github.com/azeezlala/assessment/api/internal/application_errrors"
	"github.com/azeezlala/assessment/api/internal/model"
	"github.com/azeezlala/assessment/api/internal/repository"
	"github.com/azeezlala/assessment/shared/pubsub"
	"github.com/google/uuid"
	"log"
)

type ICustomerService interface {
	CreateCustomer(ctx context.Context, data model.Customer) (*model.Customer, error)
}

type CustomerService struct {
	customer repository.ICustomerRepository
	sub      pubsub.IPubSub
}

func NewCustomerService(customer repository.ICustomerRepository, sub pubsub.IPubSub) ICustomerService {
	return &CustomerService{
		customer: customer,
		sub:      sub,
	}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, data model.Customer) (*model.Customer, error) {
	log.Printf("creating customer with email: %v", data.Email)

	if len(data.Email) < 1 {
		return nil, errors.New("email is required")
	}

	if len(data.Name) < 1 {
		return nil, errors.New("name is required")
	}

	customer, err := s.customer.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, application_errrors.ErrUnableToProcess
	}

	if customer != nil {
		return nil, errors.New("customer with the same email already exists")
	}

	data.ID = uuid.NewString()
	res, err := s.customer.CreateCustomer(ctx, &data)
	if err != nil {
		log.Printf("unable to create customer: %v", err)
		return nil, errors.New("unable to create customer")
	}

	if s.sub != nil {
		err = s.sub.Publish(pubsub.CustomerAdded, map[string]interface{}{
			data.ID: data.Email,
		})
		if err != nil {
			log.Printf("unable to publish customer added: %v", err)
		}
	}

	return res, nil
}
