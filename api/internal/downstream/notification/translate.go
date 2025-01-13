package notification

import (
	pb "github.com/azeezlala/assessment/shared/grpc/protobuf"
	"github.com/azeezlala/assessment/shared/model"
)

func ToNotification(response *pb.GetNotificationsResponse) []model.Notification {
	var notifications []model.Notification
	for _, n := range response.Notifications {
		notifications = append(notifications, model.Notification{
			ID:     n.Id,
			UserID: n.Id,
			Message: model.NotificationMessage{
				Content: n.Message,
			},
		})
	}

	return notifications
}
