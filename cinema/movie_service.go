package cinema

import (
	"time"

	"github.com/ygutara/movie-festival-app/model"
)

func (cinema Cinema) MovieGet_(id int) (record model.Movie, err error) {
	err = cinema.DB.Preload("Artists").Preload("Gendres").Where("id = ?", id).First(&record).Error
	return
}

func (cinema Cinema) MostViewedMovie_() (record model.Movie, err error) {
	query := `id = (SELECT movie.id FROM viewers
					RIGHT JOIN movie ON viewers.movie_id = movie.id
					GROUP BY movie.id
					ORDER BY COUNT(viewers.movie_id) DESC
					LIMIT 1)`

	err = cinema.DB.Preload("Artists").Preload("Gendres").Where(query).First(&record).Error
	return
}

func (cinema Cinema) MostViewedGendre_() (record model.Gendre, err error) {
	query := `id = (SELECT gendre_id FROM movie_gendre
				 	LEFT JOIN viewers ON viewers.movie_id = movie_gendre.movie_id
				 	GROUP BY gendre_id
				  	ORDER BY COUNT(viewers.movie_id) DESC
				  	LIMIT 1)`

	err = cinema.DB.Where(query).First(&record).Error
	return
}

func (cinema Cinema) MovieGetByTitle_(title string) (record model.Movie, err error) {
	err = cinema.DB.Where("title = ?", title).First(&record).Error
	return
}

func (cinema Cinema) MovieList_() (records []model.Movie, err error) {
	err = cinema.DB.Find(&records).Error
	return
}

func (cinema Cinema) MovieUpdate_(record *model.Movie) (err error) {
	existedRecord := model.Movie{}

	existedRecord, err = cinema.MovieGet_(record.ID)
	if err != nil {
		return
	}

	record.Rating = existedRecord.Rating
	record.CreatedAt = existedRecord.CreatedAt
	record.UpdatedAt = time.Now()

	cinema.DB.Exec("DELETE FROM movie_gendre WHERE movie_id = ?", record.ID)
	cinema.DB.Exec("DELETE FROM movie_artist WHERE movie_id = ?", record.ID)

	err = cinema.DB.Save(record).Error

	*record, _ = cinema.MovieGet_(record.ID)
	return
}

func (cinema Cinema) MovieCreate_(record *model.Movie) (err error) {

	record.ID = 0
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Time{}
	record.Rating = 0

	err = cinema.DB.Create(record).Error

	*record, _ = cinema.MovieGet_(record.ID)
	return
}

func (cinema Cinema) MovieDelete_(id int) (err error) {
	existedMovie := model.Movie{ID: id}
	if existedMovie, err = cinema.MovieGet_(id); err != nil {
		return
	} else {
		if existedMovie.ID == 0 {
			return model.ErrNotFound
		}
		return cinema.DB.Delete(&existedMovie).Error
	}
}
