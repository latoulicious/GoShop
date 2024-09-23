package category

import (
	"github.com/google/uuid"
	"github.com/latoulicious/GoShop/internal/product"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID       uuid.UUID         `gorm:"type:uuid"`
	Name     string            `gorm:"type:varchar(50)"`
	Products []product.Product `gorm:"foreignKey:CategoryID"`
}
