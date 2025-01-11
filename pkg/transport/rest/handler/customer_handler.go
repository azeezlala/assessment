package handler

import (
	"github.com/azeezlala/assessment/internal/repository"
	"github.com/azeezlala/assessment/internal/service"
	"github.com/azeezlala/assessment/pkg/transport/rest/model"
	"github.com/azeezlala/assessment/pkg/transport/rest/translate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CustomerHandler struct {
	customerService service.ICustomerService
}

func NewCustomerHandler() *CustomerHandler {
	return &CustomerHandler{
		customerService: service.NewCustomerService(repository.NewCustomerRepository()),
	}
}

func (h CustomerHandler) RegisterRoutes(router *gin.Engine) {
	crouter := router.Group("/customers")
	crouter.POST("/", h.createCustomer)
}

func (h *CustomerHandler) createCustomer(ctx *gin.Context) {
	var body model.CustomerRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Printf("error while parsing request body: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &model.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	res, err := h.customerService.CreateCustomer(ctx, body.ToCustomer())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &model.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, model.SuccessResponse{
		Status:  true,
		Message: "Customer created",
		Data:    translate.ToCustomerResponse(res),
	})
}
