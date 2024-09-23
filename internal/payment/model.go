package internal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid"`
	OrderID uuid.UUID `gorm:"type:uuid"`
	Amount  float64   `gorm:"type:decimal(10,2)"`
	Status  string    `gorm:"type:varchar(20)"`
}
