package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/azeezlala/assessment/internal/model"
	"github.com/azeezlala/assessment/internal/pubsub/pkg"
	"github.com/azeezlala/assessment/internal/repository"
	"log"
	"sync"
)

type NotificationService interface {
	AddNotification(ctx context.Context, userID, message string) error
	GetNotifications(ctx context.Context, userID string) ([]model.Notification, error)
	ClearNotification(ctx context.Context, userID, notificationID string) error
	ClearAllNotifications(ctx context.Context, userID string) error
}

type notification struct {
	repository repository.NotificationRepository
	once       sync.Once
}

var (
	subscriptionOnce sync.Once // Global variable to ensure the subscription happens only once
)

const (
	NotificationEvent string = "notification"
)

func NewNotificationService(repository repository.NotificationRepository, sub pkg.IPubSub) NotificationService {
	n := notification{
		repository: repository,
	}

	// dynamically subscribing to event
	if sub != nil {
		subscriptionOnce.Do(func() {
			fmt.Println("Subscribing to NotificationEvent...")
			sub.Subscribe(NotificationEvent, n.notificationJob)
		})
	}

	return &n
}

func (s *notification) AddNotification(ctx context.Context, userID, message string) error {
	err := s.repository.AddNotification(ctx, userID, message)
	if err != nil {
		log.Println("error adding notification:", err)
		return errors.New("error adding notification")
	}
	return nil
}

func (s *notification) GetNotifications(ctx context.Context, userID string) ([]model.Notification, error) {
	res, err := s.repository.GetNotifications(ctx, userID)
	if err != nil {
		log.Println("error getting notifications:", err)
		return nil, errors.New("error getting notifications")
	}

	return res, nil
}

func (s *notification) ClearNotification(ctx context.Context, userID, notificationID string) error {
	err := s.repository.ClearNotification(ctx, userID, notificationID)
	if err != nil {
		log.Println("error clearing notification:", err)
		return errors.New("error clearing notification")
	}

	return nil
}

func (s *notification) ClearAllNotifications(ctx context.Context, userID string) error {
	err := s.repository.ClearAllNotifications(ctx, userID)
	if err != nil {
		log.Println("error clearing all notifications:", err)
		return errors.New("error clearing all notifications")
	}

	return nil
}

func (s *notification) notificationJob(ctx context.Context, options pkg.Options) {
	var payload map[string]interface{}
	err := json.Unmarshal(options.Payload.([]byte), &payload)
	if err != nil {
		log.Println("error unmarshalling payload:", err)
		return
	}

	for key, value := range payload {
		log.Printf("addding notication for user: %s, message: %s", key, value)
		err := s.AddNotification(ctx, key, fmt.Sprintf("%v", value))
		if err != nil {
			log.Println("error adding notification:", err)
		}
	}
}
