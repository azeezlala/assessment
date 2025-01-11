package model

import (
	"time"
)

type Customer struct {
	ID        string    `gorm:"primary_key;type:uuid;"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;index; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
