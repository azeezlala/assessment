package service

import (
	"context"
	model2 "github.com/azeezlala/assessment/api/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockResourceRepository struct {
	mock.Mock
}

func (m *MockResourceRepository) Find(ctx context.Context) ([]model2.Resources, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]model2.Resources), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockResourceRepository) Update(ctx context.Context, resources *model2.Resources) error {
	args := m.Called(ctx, resources)
	return args.Error(0)
}

func (m *MockResourceRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockResourceRepository) FindByID(ctx context.Context, id string) (*model2.Resources, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*model2.Resources), args.Error(1)
	}
	return nil, args.Error(1)
}

type MockCustomerResourceRepository struct {
	mock.Mock
}

func (m *MockCustomerResourceRepository) FindByCustomerAndResource(ctx context.Context, customerID, resourceID string) (*model2.CustomerResource, error) {
	args := m.Called(ctx, customerID, resourceID)
	if args.Get(0) != nil {
		return args.Get(0).(*model2.CustomerResource), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCustomerResourceRepository) CreateResource(ctx context.Context, data *model2.CustomerResource) (*model2.CustomerResource, error) {
	args := m.Called(ctx, data)
	if args.Get(0) != nil {
		return args.Get(0).(*model2.CustomerResource), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCustomerResourceRepository) FindResourcesByCustomerID(ctx context.Context, id string) ([]model2.CustomerResource, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).([]model2.CustomerResource), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCustomerResourceRepository) DeleteResource(ctx context.Context, customerID, resourceID string) error {
	args := m.Called(ctx, customerID, resourceID)
	return args.Error(0)
}

// Tests
func TestCreateResource(t *testing.T) {
	mockCustomerRepo := new(MockCustomerRepository)
	mockResourceRepo := new(MockResourceRepository)
	mockCustomerResourceRepo := new(MockCustomerResourceRepository)

	service := NewCustomerResourceService(nil,
		WithCustomerRepository(mockCustomerRepo),
		WithResourceRepository(mockResourceRepo),
		WithCustomerResourceRepository(mockCustomerResourceRepo),
	)

	ctx := context.TODO()
	customerID := "1234"
	resourceID := "5678"

	t.Run("successfully create resource", func(t *testing.T) {
		mockCustomerResourceRepo.On("FindByCustomerAndResource", ctx, customerID, resourceID).Return(nil, nil)
		mockResourceRepo.On("FindByID", ctx, resourceID).Return(&model2.Resources{ID: resourceID}, nil)
		mockCustomerRepo.On("FindByID", ctx, customerID).Return(&model2.Customer{ID: customerID}, nil)
		mockCustomerResourceRepo.On("CreateResource", ctx, mock.AnythingOfType("*model.CustomerResource")).Return(&model2.CustomerResource{
			ID:         "generated-uuid",
			CustomerID: customerID,
			ResourceID: resourceID,
		}, nil)

		resource, err := service.CreateResource(ctx, model2.CustomerResource{CustomerID: customerID, ResourceID: resourceID})

		assert.NoError(t, err)
		assert.NotNil(t, resource)
		assert.Equal(t, customerID, resource.CustomerID)
		assert.Equal(t, resourceID, resource.ResourceID)
		mockCustomerResourceRepo.AssertExpectations(t)
	})
}

func TestFetchResourcesByCustomerID(t *testing.T) {
	mockCustomerResourceRepo := new(MockCustomerResourceRepository)
	service := NewCustomerResourceService(nil,
		WithCustomerResourceRepository(mockCustomerResourceRepo),
	)

	ctx := context.TODO()
	customerID := "1234"

	t.Run("successfully fetch resources", func(t *testing.T) {
		mockCustomerResourceRepo.On("FindResourcesByCustomerID", ctx, customerID).Return([]model2.CustomerResource{
			{CustomerID: customerID, ResourceID: "5678"},
		}, nil)

		resources, err := service.FetchResourcesByCustomerID(ctx, customerID)

		assert.NoError(t, err)
		assert.NotNil(t, resources)
		assert.Len(t, resources, 1)
		mockCustomerResourceRepo.AssertExpectations(t)
	})
}

func TestDeleteResource(t *testing.T) {
	mockCustomerResourceRepo := new(MockCustomerResourceRepository)
	service := NewCustomerResourceService(nil,
		WithCustomerResourceRepository(mockCustomerResourceRepo),
	)

	ctx := context.TODO()
	customerID := "1234"
	resourceID := "5678"

	t.Run("successfully delete resource", func(t *testing.T) {
		mockCustomerResourceRepo.On("DeleteResource", ctx, customerID, resourceID).Return(nil)

		err := service.DeleteResource(ctx, customerID, resourceID)

		assert.NoError(t, err)
		mockCustomerResourceRepo.AssertExpectations(t)
	})

}
