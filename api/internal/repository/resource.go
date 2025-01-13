package repository

import (
	"context"
	"errors"
	"github.com/azeezlala/assessment/api/database"
	"github.com/azeezlala/assessment/api/internal/model"
	"gorm.io/gorm"
)

type IResourceRepository interface {
	FindByID(ctx context.Context, id string) (*model.Resources, error)
	Find(ctx context.Context) ([]model.Resources, error)
	Update(context.Context, *model.Resources) error
	Delete(ctx context.Context, id string) error
}

type resourceRepository struct {
	db *gorm.DB
}

func NewResourceRepository() IResourceRepository {
	return resourceRepository{
		db: database.GetDB().DB,
	}
}

func (r resourceRepository) FindByID(ctx context.Context, id string) (*model.Resources, error) {
	var result model.Resources

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (r resourceRepository) Find(ctx context.Context) ([]model.Resources, error) {
	var result []model.Resources
	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r resourceRepository) Update(ctx context.Context, resources *model.Resources) error {
	return r.db.WithContext(ctx).Where("id = ?", resources.ID).Updates(resources).Error
}

func (r resourceRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Resources{}).Error
}
