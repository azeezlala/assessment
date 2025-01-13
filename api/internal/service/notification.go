package service

import (
	"context"
	"encoding/json"
	"errors"
	dn "github.com/azeezlala/assessment/api/internal/downstream/notification"
	"github.com/azeezlala/assessment/shared/model"
	"github.com/azeezlala/assessment/shared/pubsub"
	"log"
)

type NotificationService interface {
	GetNotifications(ctx context.Context, userID string) ([]model.Notification, error)
	ClearNotification(ctx context.Context, userID, notificationID string) error
	ClearAllNotifications(ctx context.Context, userID string) error
}

type notification struct {
	downstream *dn.Client
	sub        pubsub.IPubSub
}

func NewNotificationService(downstream *dn.Client, sub pubsub.IPubSub) NotificationService {
	n := notification{
		downstream: downstream,
		sub:        sub,
	}

	return &n
}

func (s *notification) GetNotifications(ctx context.Context, userID string) ([]model.Notification, error) {
	res, err := s.downstream.GetNotifications(ctx, userID)
	if err != nil {
		log.Println("error getting notifications:", err)
		return nil, errors.New("error getting notifications")
	}

	return res, nil
}

func (s *notification) ClearNotification(ctx context.Context, userID, notificationID string) error {
	err := s.downstream.ClearNotification(ctx, userID, notificationID)
	if err != nil {
		log.Println("error clearing notification:", err)
		return errors.New("error clearing notification")
	}

	return nil
}

func (s *notification) ClearAllNotifications(ctx context.Context, userID string) error {
	err := s.downstream.ClearAllNotifications(ctx, userID)
	if err != nil {
		log.Println("error clearing all notifications:", err)
		return errors.New("error clearing all notifications")
	}

	return nil
}

func (s *notification) notificationJob(ctx context.Context, options pubsub.Options) {
	var payload map[string]interface{}
	err := json.Unmarshal(options.Payload.([]byte), &payload)
	if err != nil {
		log.Println("error unmarshalling payload:", err)
		return
	}

	for key, value := range payload {
		log.Printf("addding notication for user: %s, message: %s", key, value)
		//err := s.AddNotification(ctx, key, fmt.Sprintf("%v", value))
		if err != nil {
			log.Println("error adding notification:", err)
		}
	}
}
