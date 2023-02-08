package auth

import (
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}
