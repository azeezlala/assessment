package router

import (
	handler2 "github.com/azeezlala/assessment/api/pkg/transport/rest/handler"
	"github.com/azeezlala/assessment/shared/pubsub"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	router *gin.Engine, job pubsub.IPubSub,
) {
	var (
		customerHandler         = handler2.NewCustomerHandler(job)
		customerResourceHandler = handler2.NewCustomerResource(job)
		notificationHandler     = handler2.NewNotificationHandler(job)
		resourceHandler         = handler2.NewResourceHandler()
	)

	customerHandler.RegisterRoutes(router)
	customerResourceHandler.RegisterRoutes(router)
	notificationHandler.RegisterRoutes(router)
	resourceHandler.RegisterRoutes(router)
}
