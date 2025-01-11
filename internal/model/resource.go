package model

type Resources struct {
	ID     string `json:"id" gorm:"primary_key;type:uuid"`
	Name   string `json:"name" gorm:"unique;index"`
	Type   string `json:"type" gorm:"index"`
	Region string `json:"region" gorm:"index"`
}
