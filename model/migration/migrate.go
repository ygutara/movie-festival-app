package migration

import (
	"github.com/ygutara/movie-festival-app/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&model.Movie{})
	db.AutoMigrate(&model.Artist{})
	db.AutoMigrate(&model.Gendre{})
	db.AutoMigrate(&model.RatingDetail{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserToken{})
	return nil
}
