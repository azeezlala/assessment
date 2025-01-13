package translate

import (
	pb "github.com/azeezlala/assessment/shared/grpc/protobuf"
	"github.com/azeezlala/assessment/shared/model"
)

func ToGetNotificationsResponse(data []model.Notification) *pb.GetNotificationsResponse {
	var result []*pb.Notification

	for _, notification := range data {
		result = append(result, &pb.Notification{
			Id:      notification.ID,
			Message: notification.Message.Content,
		})
	}

	return &pb.GetNotificationsResponse{
		Notifications: result,
	}
}
