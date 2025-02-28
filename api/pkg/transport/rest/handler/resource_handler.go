package handler

import (
	"github.com/azeezlala/assessment/api/internal/repository"
	"github.com/azeezlala/assessment/api/internal/service"
	"github.com/azeezlala/assessment/api/pkg/transport/rest/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResourceHandler struct {
	resourceService service.IResourceService
}

func NewResourceHandler() ResourceHandler {
	return ResourceHandler{
		resourceService: service.NewResourceService(repository.NewResourceRepository()),
	}
}
func (h ResourceHandler) RegisterRoutes(router *gin.Engine) {
	rroute := router.Group("/resources")
	rroute.PATCH("/", h.updateResource)
	rroute.GET("/", h.getResources)
}

func (h ResourceHandler) updateResource(c *gin.Context) {
	var body model.Resource

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, &model.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	err := h.resourceService.UpdateResource(c, body.ToResource())
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &model.SuccessResponse{
		Status:  true,
		Message: "resource updated",
	})
}

func (h ResourceHandler) getResources(c *gin.Context) {
	resources, err := h.resourceService.Find(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, &model.SuccessResponse{
		Status:  true,
		Message: "resource fetched",
		Data:    resources,
	})
}
