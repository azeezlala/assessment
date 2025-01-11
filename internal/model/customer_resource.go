package model

type CustomerResource struct {
	ID         string    `gorm:"primary_key;type:uuid;" json:"id"`
	CustomerID string    `gorm:"type:uuid;not null;index;uniqueIndex:idx_customer_resource" json:"customer_id"`
	ResourceID string    `gorm:"type:uuid;not null;index;uniqueIndex:idx_customer_resource" json:"resource_id"`
	Customer   Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CustomerID" json:"-"`
	Resources  Resources `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:ResourceID" json:"-"`
}
