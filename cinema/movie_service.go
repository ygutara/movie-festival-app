package cinema

import (
	"errors"
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

func (cinema Cinema) MovieBrowse_(search string, page int) (records []model.Movie, err error) {
	if page <= 0 {
		err = errors.New("Must be higher than 0")
	}
	search = "%" + search + "%"

	query := `SELECT 
					*, 
					ceil(count(*) OVER () / 10.0) AS max_page 
				FROM movie
				WHERE LOWER(title) LIKE LOWER(?) OR
					  LOWER(description) LIKE LOWER(?) OR
					  id IN (SELECT movie_id FROM movie_gendre
							 WHERE gendre_id IN (SELECT id FROM gendre
												 WHERE LOWER(name) LIKE LOWER(?)
												 )
							) OR
					  id IN (SELECT movie_id FROM movie_artist
							 WHERE artist_id IN (SELECT id FROM artist
												 WHERE LOWER(name) LIKE LOWER(?)
												)
							)
				ORDER BY movie.id
				LIMIT 10 OFFSET (? - 1) * 10`

	err = cinema.DB.Raw(query, search, search, search, search, page).Scan(&records).Error
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
