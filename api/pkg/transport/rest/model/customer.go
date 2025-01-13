package model

import (
	"github.com/azeezlala/assessment/api/internal/model"
)

type (
	CustomerRequest struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}

	CustomerResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func (c CustomerRequest) ToCustomer() model.Customer {
	return model.Customer{
		Name:  c.Name,
		Email: c.Email,
	}
}
