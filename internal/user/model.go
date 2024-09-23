package internal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Username string    `gorm:"type:varchar(20)"`
	Password string    `gorm:"type:varchar(60)"`
	Name     string    `gorm:"type:varchar(50)"`
}
