package repository

import (
	"context"
	"github.com/azeezlala/assessment/database"
	"github.com/azeezlala/assessment/internal/model"
	"gorm.io/gorm"
)

type ICustomerRepository interface {
	CreateCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, error)
	FindByID(ctx context.Context, id string) (*model.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository() ICustomerRepository {
	return customerRepository{
		db: database.GetDB().DB,
	}
}

func (c customerRepository) CreateCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	err := c.db.WithContext(ctx).Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c customerRepository) FindByID(ctx context.Context, id string) (*model.Customer, error) {
	var result model.Customer

	if err := c.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
