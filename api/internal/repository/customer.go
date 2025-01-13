package repository

import (
	"context"
	"github.com/azeezlala/assessment/api/database"
	"github.com/azeezlala/assessment/api/internal/model"
	"gorm.io/gorm"
)

type ICustomerRepository interface {
	CreateCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, error)
	FindByID(ctx context.Context, id string) (*model.Customer, error)
	FindByEmail(ctx context.Context, email string) (*model.Customer, error)
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

func (c customerRepository) FindByEmail(ctx context.Context, email string) (*model.Customer, error) {
	var result model.Customer
	if err := c.db.WithContext(ctx).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
