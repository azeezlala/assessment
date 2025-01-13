package model

import (
	"github.com/azeezlala/assessment/api/internal/model"
)

type CustomerResourceRequest struct {
	CustomerID string `json:"customer_id" binding:"required"`
	ResourceID string `json:"resource_id" binding:"required"`
}

func (c CustomerResourceRequest) ToCustomerResource() model.CustomerResource {
	return model.CustomerResource{
		CustomerID: c.CustomerID,
		ResourceID: c.ResourceID,
	}
}
