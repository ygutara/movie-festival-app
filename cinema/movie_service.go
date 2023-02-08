package cinema

import (
	"time"

	"github.com/ygutara/movie-festival-app/model"
)

func (cinema Cinema) MovieGet_(id int) (record model.Movie, err error) {
	err = cinema.DB.Where("id = ?", id).First(&record).Error
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

	record.CreatedAt = existedRecord.CreatedAt
	record.UpdatedAt = time.Now()

	err = cinema.DB.Save(record).Error
	return
}

func (cinema Cinema) MovieCreate_(record *model.Movie) (err error) {

	record.ID = 0
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Time{}

	err = cinema.DB.Create(record).Error

	return
}

func (cinema Cinema) MovieDelete_(id int) (err error) {
	existedMovie := model.Movie{ID: id}
	if existedMovie, err = cinema.MovieGet_(id); err != nil {
		return
	} else {
		if existedMovie == (model.Movie{}) {
			return model.ErrNotFound
		}
		return cinema.DB.Delete(&existedMovie).Error
	}
}
