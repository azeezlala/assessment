package handler

import (
	repository2 "github.com/azeezlala/assessment/api/internal/repository"
	"github.com/azeezlala/assessment/api/internal/service"
	model2 "github.com/azeezlala/assessment/api/pkg/transport/rest/model"
	"github.com/azeezlala/assessment/shared/pubsub"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CustomerResourceHandler struct {
	customerResourceService service.ICustomerResourceService
}

func NewCustomerResource(sub pubsub.IPubSub) *CustomerResourceHandler {
	customerRepository := service.NewCustomerResourceService(sub,
		service.WithCustomerRepository(repository2.NewCustomerRepository()),
		service.WithResourceRepository(repository2.NewResourceRepository()),
		service.WithCustomerResourceRepository(repository2.NewCustomerResourceRepository()),
	)
	return &CustomerResourceHandler{customerResourceService: customerRepository}
}
func (h CustomerResourceHandler) RegisterRoutes(router *gin.Engine) {
	crouter := router.Group("/customer-resources")
	crouter.POST("/", h.createCustomerResource)
	crouter.GET("/:customerId", h.getCustomerResource)
	crouter.DELETE("/", h.deleteResource)
}

func (h *CustomerResourceHandler) createCustomerResource(ctx *gin.Context) {
	var body model2.CustomerResourceRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Printf("error while parsing request body: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &model2.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	res, err := h.customerResourceService.CreateResource(ctx, body.ToCustomerResource())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &model2.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &model2.SuccessResponse{
		Status:  true,
		Message: "resource created",
		Data:    res,
	})
}

func (h *CustomerResourceHandler) getCustomerResource(ctx *gin.Context) {
	customerID := ctx.Param("customerId")

	res, err := h.customerResourceService.FetchResourcesByCustomerID(ctx, customerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &model2.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &model2.SuccessResponse{
		Status:  true,
		Message: "fetched resource successfully",
		Data:    res,
	})
}

func (h *CustomerResourceHandler) deleteResource(ctx *gin.Context) {
	var body model2.CustomerResourceRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Printf("error while parsing request body: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &model2.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	err := h.customerResourceService.DeleteResource(ctx, body.ResourceID, body.CustomerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &model2.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
