package model

import (
	"time"
)

type Movie struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Duration      float64        `json:"duration"`
	Image         string         `json:"image"`
	URL           string         `json:"url"`
	Artists       []Artist       `json:"artists" gorm:"many2many:movie_artist;save_associations:false"`
	Gendres       []Gendre       `json:"gendres" gorm:"many2many:movie_gendre;save_associations:false"`
	Viewers       []User         `json:"viewers" gorm:"many2many:viewers;save_associations:false"`
	Rating        float64        `json:"rating"`
	RatingDetails []RatingDetail `json:"rating_details"`
	MaxPage       int            `json:"-"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (Movie) TableName() string {
	return "movie"
}

type Artist struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (Artist) TableName() string {
	return "artist"
}

type Gendre struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (Gendre) TableName() string {
	return "gendre"
}

type RatingDetail struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	MovieID int     `json:"movie_id"`
	Movie   *Movie  `json:"movie" gorm:"save_associations:false"`
	UserID  int     `json:"user_id"`
	User    *User   `json:"user" gorm:"save_associations:false"`
	Rating  float64 `json:"rating"`
}

func (RatingDetail) TableName() string {
	return "rating_detail"
}
