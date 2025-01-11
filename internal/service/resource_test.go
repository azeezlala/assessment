package service

import (
	"context"
	"errors"
	"github.com/azeezlala/assessment/internal/application_errrors"
	"github.com/azeezlala/assessment/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceService(t *testing.T) {
	mockRepo := new(MockResourceRepository)
	service := NewResourceService(mockRepo)
	ctx := context.TODO()

	t.Run("UpdateResource - Success", func(t *testing.T) {
		resource := model.Resources{ID: "1", Name: "Test Resource"}
		mockRepo.On("FindByID", ctx, "1").Return(&resource, nil)
		mockRepo.On("Update", ctx, &resource).Return(nil)

		err := service.UpdateResource(ctx, resource)
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateResource - Resource Not Found", func(t *testing.T) {
		resource := model.Resources{ID: "2", Name: "Nonexistent Resource"}
		mockRepo.On("FindByID", ctx, "2").Return(nil, nil)

		err := service.UpdateResource(ctx, resource)
		assert.Error(t, err)
		assert.Equal(t, "resource not found", err.Error())

		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateResource - FindByID Error", func(t *testing.T) {
		resource := model.Resources{ID: "3", Name: "Error Resource"}
		mockRepo.On("FindByID", ctx, "3").Return(nil, errors.New("db error"))

		err := service.UpdateResource(ctx, resource)
		assert.Error(t, err)
		assert.Equal(t, application_errrors.ErrUnableToProcess, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateResource - Update Error", func(t *testing.T) {
		resource := model.Resources{ID: "4", Name: "Update Error Resource"}
		mockRepo.On("FindByID", ctx, "4").Return(&resource, nil)
		mockRepo.On("Update", ctx, &resource).Return(errors.New("db error"))

		err := service.UpdateResource(ctx, resource)
		assert.Error(t, err)
		assert.Equal(t, application_errrors.ErrUnableToProcess, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Find - Success", func(t *testing.T) {
		expectedResources := []model.Resources{
			{ID: "1", Name: "Resource 1"},
			{ID: "2", Name: "Resource 2"},
		}
		mockRepo.On("Find", ctx).Return(expectedResources, nil)

		resources, err := service.Find(ctx)
		assert.NoError(t, err)
		assert.Equal(t, expectedResources, resources)

		mockRepo.AssertExpectations(t)
	})

}
