package main

import (
	"fmt"

	"github.com/ygutara/movie-festival-app/app/config"
	"github.com/ygutara/movie-festival-app/model"
	"github.com/ygutara/movie-festival-app/model/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("MIGRATION failed")
		panic("error connect to DB")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	migration.Migrate(db)

	CreateUserAdmin(db)
	CreateArtistData(db)
	CreateGendreData(db)

	fmt.Println("MIGRATION SUCCESS")
}

func CreateUserAdmin(db *gorm.DB) {
	user := model.User{}
	user.Username = "admin"
	user.Password = "admin"
	user.IsAdmin = true

	userExisted := model.User{}
	db.Where("username = ?", user.Username).Find(&userExisted)

	if userExisted.Username != user.Username {
		db.Create(&user)
	}
}

func CreateArtistData(db *gorm.DB) {
	artists := []model.Artist{
		{Name: "Meryl Streep"},
		{Name: "Tom Cruise"},
		{Name: "Angelina Jolie"},
		{Name: "Brad Pitt"},
		{Name: "Leonardo DiCaprio"},
		{Name: "Sandra Bullock"},
		{Name: "Johnny Depp"},
		{Name: "Jennifer Lawrence"},
		{Name: "Will Smith"},
		{Name: "Denzel Washington"},
	}

	for i := range artists {
		artist := model.Artist{}
		db.Where("name = ?", artists[i].Name).Find(&artist)
		if artist.Name != artists[i].Name {
			db.Create(&artists[i])
		}
	}
}

func CreateGendreData(db *gorm.DB) {
	gendres := []model.Gendre{
		{Name: "Action"},
		{Name: "Romance"},
		{Name: "Comedy"},
		{Name: "Drama"},
		{Name: "ScienceFiction"},
		{Name: "Horror"},
		{Name: "Adventure"},
		{Name: "Thriller"},
		{Name: "Fantasy"},
		{Name: "Crime/Mystery"},
	}

	for i := range gendres {
		gendre := model.Gendre{}
		db.Where("name = ?", gendres[i].Name).Find(&gendre)
		if gendre.Name != gendres[i].Name {
			db.Create(&gendres[i])
		}
	}
}
