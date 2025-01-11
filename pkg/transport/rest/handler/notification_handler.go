package handler

import (
	"fmt"
	"github.com/azeezlala/assessment/internal/pubsub/pkg"
	"github.com/azeezlala/assessment/internal/repository"
	"github.com/azeezlala/assessment/internal/service"
	"github.com/azeezlala/assessment/pkg/transport/rest/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationHandler struct {
	NotificationService service.NotificationService
}

func NewNotificationHandler(sub pkg.IPubSub) *NotificationHandler {
	fmt.Println("i'm here")
	return &NotificationHandler{
		NotificationService: service.NewNotificationService(repository.NewNotificationRepository(), sub),
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
