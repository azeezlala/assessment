package model

import (
	"github.com/azeezlala/assessment/api/internal/model"
)

type Resource struct {
	Name   string `json:"name"  binding:"required"`
	Type   string `json:"type"  binding:"required"`
	Region string `json:"region"  binding:"required"`
}

func (r Resource) ToResource() model.Resources {
	return model.Resources{
		Name:   r.Name,
		Type:   r.Type,
		Region: r.Region,
	}
}
