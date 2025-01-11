package service

import (
	"context"
	"errors"
	"github.com/azeezlala/assessment/internal/application_errrors"
	"github.com/azeezlala/assessment/internal/model"
	"github.com/azeezlala/assessment/internal/repository"
	"log"
)

type IResourceService interface {
	UpdateResource(ctx context.Context, data model.Resources) error
	Find(ctx context.Context) ([]model.Resources, error)
}

type resource struct {
	resourceRepository repository.IResourceRepository
}

func NewResourceService(resourceRepository repository.IResourceRepository) IResourceService {
	return resource{resourceRepository: resourceRepository}
}

func (r resource) UpdateResource(ctx context.Context, data model.Resources) error {
	res, err := r.resourceRepository.FindByID(ctx, data.ID)
	if err != nil {
		return application_errrors.ErrUnableToProcess
	}

	if res == nil {
		return errors.New("resource not found")
	}

	if err = r.resourceRepository.Update(ctx, &data); err != nil {
		log.Printf("error updating resource: %v", err)
		return application_errrors.ErrUnableToProcess
	}

	return nil
}

func (r resource) Find(ctx context.Context) ([]model.Resources, error) {
	res, err := r.resourceRepository.Find(ctx)
	if err != nil {
		return []model.Resources{}, err
	}

	return res, nil
}
