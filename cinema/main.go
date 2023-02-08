package cinema

import (
	"gorm.io/gorm"
)

type Cinema struct {
	DB *gorm.DB
}
