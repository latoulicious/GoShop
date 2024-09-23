package cart

import (
	"github.com/google/uuid"
	"github.com/latoulicious/GoShop/internal/cartproduct"
	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model
	ID         uuid.UUID                 `gorm:"type:uuid"`
	CustomerID uuid.UUID                 `gorm:"type:uuid"`
	Products   []cartproduct.CartProduct `gorm:"many2many:shopping_cart_products;"`
}
