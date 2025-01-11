package router

import (
	"github.com/azeezlala/assessment/internal/pubsub/pkg"
	"github.com/azeezlala/assessment/pkg/transport/rest/handler"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	router *gin.Engine, job pkg.IPubSub,
) {
	var (
		customerHandler         = handler.NewCustomerHandler()
		customerResourceHandler = handler.NewCustomerResource()
		notificationHandler     = handler.NewNotificationHandler(job)
		resourceHandler         = handler.NewResourceHandler()
	)

	customerHandler.RegisterRoutes(router)
	customerResourceHandler.RegisterRoutes(router)
	notificationHandler.RegisterRoutes(router)
	resourceHandler.RegisterRoutes(router)
}
