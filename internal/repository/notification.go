package repository

import (
	"context"
	"errors"
	"github.com/azeezlala/assessment/internal/model"
	"github.com/google/uuid"
	"sync"
)

type NotificationRepository interface {
	AddNotification(ctx context.Context, userID, message string) error
	GetNotifications(ctx context.Context, userID string) ([]model.Notification, error)
	ClearNotification(ctx context.Context, userID, notificationID string) error
	ClearAllNotifications(ctx context.Context, userID string) error
}

type NotificationObj struct {
	mu            sync.RWMutex
	notifications map[string][]model.Notification // userID -> notifications
}

var (
	instance *NotificationObj
	once     sync.Once
)

func NewNotificationRepository() NotificationRepository {
	once.Do(func() {
		instance = &NotificationObj{
			notifications: make(map[string][]model.Notification),
		}
	})
	return instance
}

func (s *NotificationObj) AddNotification(ctx context.Context, userID, message string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	notification := model.Notification{
		ID:      uuid.NewString(),
		UserID:  userID,
		Message: message,
	}

	s.notifications[userID] = append(s.notifications[userID], notification)
	return nil
}

func (s *NotificationObj) GetNotifications(ctx context.Context, userID string) ([]model.Notification, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.notifications[userID], nil
}

func (s *NotificationObj) ClearNotification(ctx context.Context, userID, notificationID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	notifications := s.notifications[userID]
	for i, n := range notifications {
		if n.ID == notificationID {
			s.notifications[userID] = append(notifications[:i], notifications[i+1:]...)
			return nil
		}
	}
	return errors.New("notification not found")
}

func (s *NotificationObj) ClearAllNotifications(ctx context.Context, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.notifications, userID)
	return nil
}
