package model

import (
	"github.com/azeezlala/assessment/api/database"
)

func init() {
	database.RegisterModel(
		Customer{},
		Resources{},
		CustomerResource{})
}
