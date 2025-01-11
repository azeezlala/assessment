package service

import (
	"context"
	"testing"

	"github.com/azeezlala/assessment/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) FindByID(ctx context.Context, id string) (*model.Customer, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCustomerRepository) CreateCustomer(ctx context.Context, data *model.Customer) (*model.Customer, error) {
	args := m.Called(ctx, data)
	return args.Get(0).(*model.Customer), args.Error(1)
}

func TestCreateCustomer(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	customerService := NewCustomerService(mockRepo)

	t.Run("successfully create customer", func(t *testing.T) {
		// Arrange
		ctx := context.TODO()
		customer := model.Customer{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}

		mockRepo.On("CreateCustomer", ctx, mock.Anything).Return(&model.Customer{
			ID:    uuid.NewString(),
			Name:  customer.Name,
			Email: customer.Email,
		}, nil)

		// Act
		result, err := customerService.CreateCustomer(ctx, customer)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, customer.Name, result.Name)
		assert.Equal(t, customer.Email, result.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("fail when email is missing", func(t *testing.T) {
		// Arrange
		ctx := context.TODO()
		customer := model.Customer{
			Name: "John Doe",
		}

		// Act
		result, err := customerService.CreateCustomer(ctx, customer)

		// Assert
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Equal(t, "email is required", err.Error())
	})

	t.Run("fail when name is missing", func(t *testing.T) {
		// Arrange
		ctx := context.TODO()
		customer := model.Customer{
			Email: "john.doe@example.com",
		}

		// Act
		result, err := customerService.CreateCustomer(ctx, customer)

		// Assert
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Equal(t, "name is required", err.Error())
	})
}
