package cinema

import (
	"github.com/ygutara/movie-festival-app/auth"
	"gorm.io/gorm"
)

type Cinema struct {
	DB            *gorm.DB
	Authorization *auth.Auth
}
