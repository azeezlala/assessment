package repository

import (
	"context"
	"errors"
	"github.com/azeezlala/assessment/database"
	"github.com/azeezlala/assessment/internal/model"
	"gorm.io/gorm"
)

type ICustomerResourceRepository interface {
	CreateResource(ctx context.Context, data *model.CustomerResource) (*model.CustomerResource, error)
	FindByCustomerAndResource(ctx context.Context, customerID, resourceID string) (*model.CustomerResource, error)
	FindResourcesByCustomerID(ctx context.Context, customerID string) ([]model.CustomerResource, error)
	DeleteResource(ctx context.Context, customerID, resourceID string) error
}

type CustomerResourceRepository struct {
	db *gorm.DB
}

func NewCustomerResourceRepository() ICustomerResourceRepository {
	return CustomerResourceRepository{
		db: database.GetDB().DB,
	}
}

func (c CustomerResourceRepository) CreateResource(ctx context.Context, data *model.CustomerResource) (*model.CustomerResource, error) {
	if err := c.db.WithContext(ctx).Create(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c CustomerResourceRepository) FindByCustomerAndResource(ctx context.Context, customerID, resourceID string) (*model.CustomerResource, error) {
	var result model.CustomerResource

	if err := c.db.Where("customer_id = ? AND resource_id = ?", customerID, resourceID).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	return &result, nil
}

func (c CustomerResourceRepository) FindResourcesByCustomerID(ctx context.Context, customerID string) ([]model.CustomerResource, error) {
	var result = make([]model.CustomerResource, 0)

	if err := c.db.WithContext(ctx).Where("customer_id = ?", customerID).Preload("Customer").Preload("Resources").Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (c CustomerResourceRepository) DeleteResource(ctx context.Context, customerID, resourceID string) error {
	return c.db.WithContext(ctx).Where("customer_id = ? AND resource_id = ?", customerID, resourceID).Delete(&model.CustomerResource{}).Error
}
