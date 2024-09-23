package cartproduct

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartProduct struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid"`
	ShoppingCartID uuid.UUID `gorm:"type:uuid"`
	ProductID      uuid.UUID `gorm:"type:uuid"`
	Quantity       int       `gorm:"type:int"`
}
