package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/ygutara/movie-festival-app/app/config"
	"github.com/ygutara/movie-festival-app/auth"
	"github.com/ygutara/movie-festival-app/cinema"
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
		fmt.Println(err.Error())
		panic("error connect to DB")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := mux.NewRouter()

	auth_ := auth.Auth{DB: db}
	auth_.AuthRoute(router)

	cinema_ := cinema.Cinema{DB: db, Authorization: &auth_}
	cinema_.MovieRoute(router)

	// CORS
	corsWrapper := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "POST", "DELETE", "PATCH"},
		AllowCredentials: true,
	})

	// Handler
	fmt.Println("Listening to port " + config.APP_PORT + "...")
	handler := corsWrapper.Handler(router)
	http.ListenAndServe(":"+config.APP_PORT, handler)
}
