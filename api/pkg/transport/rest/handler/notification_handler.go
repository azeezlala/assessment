package handler

import (
	dn "github.com/azeezlala/assessment/api/internal/downstream/notification"
	"github.com/azeezlala/assessment/api/internal/service"
	"github.com/azeezlala/assessment/api/pkg/transport/rest/model"
	"github.com/azeezlala/assessment/shared/pubsub"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type NotificationHandler struct {
	NotificationService service.NotificationService
}

func NewNotificationHandler(sub pubsub.IPubSub) *NotificationHandler {
	downstream, err := dn.New()
	if err != nil {
		log.Fatal(err)
	}

	return &NotificationHandler{
		NotificationService: service.NewNotificationService(downstream, sub),
	}
}

func (h NotificationHandler) RegisterRoutes(router *gin.Engine) {
	nrouter := router.Group("/notifications")
	nrouter.GET("/:userId", h.getNotification)
	nrouter.DELETE("/:userId", h.deleteNotification)
}

func (h NotificationHandler) getNotification(c *gin.Context) {
	userID := c.Param("userId")

	res, err := h.NotificationService.GetNotifications(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: true,
		Data:   res,
	})
}

func (h NotificationHandler) deleteNotification(c *gin.Context) {
	userID := c.Param("userId")

	err := h.NotificationService.ClearAllNotifications(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &model.SuccessResponse{
		Status: true,
	})
}
