package model

import "github.com/azeezlala/assessment/database"

func init() {
	database.RegisterModel(
		Customer{},
		Resources{},
		CustomerResource{})
}
