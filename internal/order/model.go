package internal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid"`
	CustomerID     uuid.UUID `gorm:"type:uuid"`
	ShoppingCartID uuid.UUID `gorm:"type:uuid"`
	TotalPrice     float64   `gorm:"type:decimal(10,2)"`
	Status         string    `gorm:"type:varchar(20)"`
}
