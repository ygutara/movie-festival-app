package migration

import (
	"github.com/ygutara/movie-festival-app/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&model.Movie{})

	return nil
}
