package model

import (
	"time"
)

type Movie struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    float64   `json:"duration"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Movie) TableName() string {
	return "movie"
}
