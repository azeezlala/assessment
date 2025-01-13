package notification

import (
	"context"
	pb "github.com/azeezlala/assessment/shared/grpc/protobuf"
	"github.com/azeezlala/assessment/shared/model"
)

func (c *Client) GetNotifications(ctx context.Context, userID string) ([]model.Notification, error) {
	response, err := c.client.GetNotifications(ctx, &pb.GetNotificationsRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	return ToNotification(response), nil
}

func (c *Client) ClearNotification(ctx context.Context, userID, notificationID string) error {
	_, err := c.client.ClearNotification(ctx, &pb.ClearNotificationRequest{
		UserId:         userID,
		NotificationId: notificationID,
	})

	return err
}

func (c *Client) ClearAllNotifications(ctx context.Context, userID string) error {
	_, err := c.client.ClearAllNotifications(ctx, &pb.ClearAllNotificationsRequest{
		UserId: userID,
	})

	return err
}
