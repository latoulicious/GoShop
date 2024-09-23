package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid"`
	Name        string    `gorm:"type:varchar(100)"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:decimal(10,2)"`
	CategoryID  uuid.UUID `gorm:"type:uuid"`
}
