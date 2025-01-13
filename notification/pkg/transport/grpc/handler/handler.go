package handler

import (
	"context"
	"github.com/azeezlala/assessment/notification/internal/repository"
	"github.com/azeezlala/assessment/notification/internal/service"
	"github.com/azeezlala/assessment/notification/pkg/transport/grpc/translate"
	pb "github.com/azeezlala/assessment/shared/grpc/protobuf"
	pubsub "github.com/azeezlala/assessment/shared/pubsub"
	"google.golang.org/grpc"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer
	notificationService service.NotificationService
}

func NewNotificationServer(s grpc.ServiceRegistrar, sub pubsub.IPubSub) {
	notification := NotificationServer{
		notificationService: service.NewNotificationService(repository.NewNotificationRepository(), sub),
	}
	pb.RegisterNotificationServiceServer(s, notification)
}

func (n NotificationServer) GetNotifications(ctx context.Context, request *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	res, err := n.notificationService.GetNotifications(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	return translate.ToGetNotificationsResponse(res), nil
}

func (n NotificationServer) ClearNotification(ctx context.Context, request *pb.ClearNotificationRequest) (*pb.ClearNotificationResponse, error) {
	if err := n.notificationService.ClearNotification(ctx, request.UserId, request.NotificationId); err != nil {
		return nil, err
	}

	return &pb.ClearNotificationResponse{
		Success: true,
	}, nil
}

func (n NotificationServer) ClearAllNotifications(ctx context.Context, request *pb.ClearAllNotificationsRequest) (*pb.ClearAllNotificationsResponse, error) {
	if err := n.notificationService.ClearAllNotifications(ctx, request.UserId); err != nil {
		return nil, err
	}

	return &pb.ClearAllNotificationsResponse{
		Success: true,
	}, nil
}
