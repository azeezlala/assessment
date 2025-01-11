package service

import (
	"context"
	"fmt"
	"github.com/azeezlala/assessment/internal/application_errrors"
	"github.com/azeezlala/assessment/internal/model"
	"github.com/azeezlala/assessment/internal/repository"
	"github.com/google/uuid"
	"log"
)

type ICustomerResourceService interface {
	CreateResource(ctx context.Context, data model.CustomerResource) (*model.CustomerResource, error)
	FetchResourcesByCustomerID(ctx context.Context, id string) ([]model.CustomerResource, error)
	DeleteResource(ctx context.Context, resourceID, customerID string) error
}

type customerResourceService struct {
	customer         repository.ICustomerRepository
	resource         repository.IResourceRepository
	customerResource repository.ICustomerResourceRepository
}

type CustomerResourceServiceOption func(bs *customerResourceService)

func NewCustomerResourceService(opts ...CustomerResourceServiceOption) ICustomerResourceService {
	c := &customerResourceService{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (s customerResourceService) CreateResource(ctx context.Context, data model.CustomerResource) (*model.CustomerResource, error) {
	if len(data.ResourceID) < 1 {
		return nil, fmt.Errorf("resource id is empty")
	}

	if len(data.CustomerID) < 1 {
		return nil, fmt.Errorf("customer id is empty")
	}

	res, err := s.customerResource.FindByCustomerAndResource(ctx, data.CustomerID, data.ResourceID)
	if err != nil {
		log.Printf("database error while getting customer resource: %v", err)
		return nil, application_errrors.ErrUnableToProcess
	}

	if res != nil {
		return nil, fmt.Errorf("customer resource already exists")
	}

	resource, err := s.resource.FindByID(ctx, data.ResourceID)
	if err != nil {
		log.Printf("error finding resource: %v", err)
		return nil, application_errrors.ErrUnableToProcess
	}

	if resource == nil {
		return nil, fmt.Errorf("resource not found")
	}

	customer, err := s.customer.FindByID(ctx, data.CustomerID)
	if err != nil {
		log.Printf("error finding customer: %v", err)
		return nil, application_errrors.ErrUnableToProcess
	}

	if customer == nil {
		log.Printf("customer not found")
		return nil, fmt.Errorf("customer not found")
	}

	res, err = s.customerResource.CreateResource(ctx, &model.CustomerResource{ID: uuid.NewString(), CustomerID: data.CustomerID, ResourceID: data.ResourceID})
	if err != nil {
		log.Printf("error creating customer resource: %v", err)
		return nil, application_errrors.ErrUnableToProcess
	}

	return res, nil
}

func (s customerResourceService) FetchResourcesByCustomerID(ctx context.Context, id string) ([]model.CustomerResource, error) {
	res, err := s.customerResource.FindResourcesByCustomerID(ctx, id)
	if err != nil {
		log.Printf("error finding customer resources: %v", err)
		return nil, application_errrors.ErrUnableToProcess
	}

	return res, nil
}

func (s customerResourceService) DeleteResource(ctx context.Context, customerID, resourceID string) error {
	if err := s.customerResource.DeleteResource(ctx, customerID, resourceID); err != nil {
		log.Printf("error deleting customer resource: %v", err)
		return application_errrors.ErrUnableToProcess
	}

	return nil
}

func WithCustomerRepository(customer repository.ICustomerRepository) CustomerResourceServiceOption {
	return func(bs *customerResourceService) {
		bs.customer = customer
	}
}

func WithCustomerResourceRepository(cr repository.ICustomerResourceRepository) CustomerResourceServiceOption {
	return func(bs *customerResourceService) {
		bs.customerResource = cr
	}
}

func WithResourceRepository(r repository.IResourceRepository) CustomerResourceServiceOption {
	return func(bs *customerResourceService) {
		bs.resource = r
	}
}
