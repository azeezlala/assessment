package translate

import (
	"github.com/azeezlala/assessment/internal/model"
	pb "github.com/azeezlala/assessment/pkg/transport/grpc/protobuf"
)

func ToGetNotificationsResponse(data []model.Notification) *pb.GetNotificationsResponse {
	var result []*pb.Notification

	for _, notification := range data {
		result = append(result, &pb.Notification{
			Id:      notification.ID,
			Message: notification.Message,
		})
	}

	return &pb.GetNotificationsResponse{
		Notifications: result,
	}
}
