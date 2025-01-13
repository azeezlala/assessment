package handler

import (
	"github.com/azeezlala/assessment/api/internal/repository"
	"github.com/azeezlala/assessment/api/internal/service"
	model2 "github.com/azeezlala/assessment/api/pkg/transport/rest/model"
	"github.com/azeezlala/assessment/api/pkg/transport/rest/translate"
	"github.com/azeezlala/assessment/shared/pubsub"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CustomerHandler struct {
	customerService service.ICustomerService
}

func NewCustomerHandler(sub pubsub.IPubSub) *CustomerHandler {
	return &CustomerHandler{
		customerService: service.NewCustomerService(repository.NewCustomerRepository(), sub),
	}
}

func (h CustomerHandler) RegisterRoutes(router *gin.Engine) {
	crouter := router.Group("/customers")
	crouter.POST("/", h.createCustomer)
}

func (h *CustomerHandler) createCustomer(ctx *gin.Context) {
	var body model2.CustomerRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Printf("error while parsing request body: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &model2.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	res, err := h.customerService.CreateCustomer(ctx, body.ToCustomer())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &model2.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, model2.SuccessResponse{
		Status:  true,
		Message: "Customer created",
		Data:    translate.ToCustomerResponse(res),
	})
}
