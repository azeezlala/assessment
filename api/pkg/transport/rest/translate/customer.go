package translate

import (
	"github.com/azeezlala/assessment/api/internal/model"
	apiModel "github.com/azeezlala/assessment/api/pkg/transport/rest/model"
)

func ToCustomerResponse(customer *model.Customer) apiModel.CustomerResponse {
	return apiModel.CustomerResponse{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
	}
}
