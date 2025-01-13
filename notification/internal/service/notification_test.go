package service

import (
	"context"
	"github.com/azeezlala/assessment/shared/model"
	"github.com/azeezlala/assessment/shared/pubsub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockNotificationRepository struct {
	mock.Mock
}

func (m *MockNotificationRepository) AddNotification(ctx context.Context, userID, message string) error {
	args := m.Called(ctx, userID, message)
	return args.Error(0)
}

func (m *MockNotificationRepository) GetNotifications(ctx context.Context, userID string) ([]model.Notification, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]model.Notification), args.Error(1)
}

func (m *MockNotificationRepository) ClearNotification(ctx context.Context, userID, notificationID string) error {
	args := m.Called(ctx, userID, notificationID)
	return args.Error(0)
}

func (m *MockNotificationRepository) ClearAllNotifications(ctx context.Context, userID string) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

type MockPubSub struct {
	mock.Mock
}

func (m *MockPubSub) Subscribe(event string, handler pubsub.Handler) {
	m.Called(event, handler)
}

func (m *MockPubSub) Publish(event string, payload interface{}) error {
	m.Called(event, payload)
	return nil
}

func TestNotificationService(t *testing.T) {
	mockRepo := new(MockNotificationRepository)
	mockPubSub := new(MockPubSub)
	mockPubSub.On("Subscribe", NotificationEvent, mock.Anything).Return()
	ctx := context.TODO()

	service := NewNotificationService(mockRepo, mockPubSub)

	t.Run("AddNotification Success", func(t *testing.T) {
		mockRepo.On("AddNotification", ctx, "user1", "Test Message").Return(nil)

		err := service.AddNotification(ctx, "user1", "Test Message")
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("GetNotifications Success", func(t *testing.T) {
		expectedNotifications := []model.Notification{
			{ID: "1", UserID: "user1", Message: "Message 1"},
			{ID: "2", UserID: "user1", Message: "Message 2"},
		}

		mockRepo.On("GetNotifications", ctx, "user1").Return(expectedNotifications, nil)

		notifications, err := service.GetNotifications(ctx, "user1")
		assert.NoError(t, err)
		assert.Equal(t, expectedNotifications, notifications)

		mockRepo.AssertExpectations(t)
	})

	t.Run("ClearNotification Success", func(t *testing.T) {
		mockRepo.On("ClearNotification", ctx, "user1", "notif1").Return(nil)

		err := service.ClearNotification(ctx, "user1", "notif1")
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("ClearAllNotifications Success", func(t *testing.T) {
		mockRepo.On("ClearAllNotifications", ctx, "user1").Return(nil)

		err := service.ClearAllNotifications(ctx, "user1")
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})
}
