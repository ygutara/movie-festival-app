package cinema

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ygutara/movie-festival-app/helper"
	"github.com/ygutara/movie-festival-app/model"
)

func (cinema Cinema) MovieRoute(r *mux.Router) {
	r.HandleFunc("/admin/movie/create", cinema.MovieCreate).Methods("POST")
	r.HandleFunc("/admin/movie/update", cinema.MovieUpdate).Methods("PATCH")
	r.HandleFunc("/admin/movie/most_viewed_movie", cinema.MostViewedMovie).Methods("GET")
	r.HandleFunc("/admin/movie/most_viewed_gendre", cinema.MostViewedGendre).Methods("GET")

	r.HandleFunc("/admin/movie/{id:[0-9]+}", cinema.MovieDelete).Methods("DELETE")

	r.HandleFunc("/user/movie/browse", cinema.MovieBrowse).Methods("GET")
	r.HandleFunc("/user/movie/{id:[0-9]+}", cinema.MovieGet).Methods("GET")
	r.HandleFunc("/user/movie/watch/{id:[0-9]+}", cinema.MovieWatch).Methods("GET")
	r.HandleFunc("/user/movie/vote", cinema.MovieVote).Methods("POST")
	r.HandleFunc("/user/movie/unvote", cinema.MovieUnvote).Methods("POST")
	r.HandleFunc("/user/movie/list_voted", cinema.MovieListVoted).Methods("GET")

	r.HandleFunc("/user/movie/{title:[a-zA-Z0-9-_]+}", cinema.MovieGetByTitle).Methods("GET")
}

func (cinema Cinema) MovieCreate(w http.ResponseWriter, r *http.Request) {
	if authorization := cinema.Authorization.Authorization(r); authorization == nil || !authorization.IsAdmin {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	record := model.Movie{}
	if err := decoder.Decode(&record); err == nil {
		if err := cinema.MovieCreate_(&record); err == nil {
			helper.WriteResponse(w, http.StatusOK, record, nil)
		} else {
			helper.WriteResponse(w, http.StatusInternalServerError, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}
}

func (cinema Cinema) MovieBrowse(w http.ResponseWriter, r *http.Request) {
	if authorization := cinema.Authorization.Authorization(r); authorization == nil {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	param := struct {
		Search string `json:"search"`
		Page   int    `json:"page"`
	}{}

	if err := decoder.Decode(&param); err == nil {
		if records, err := cinema.MovieBrowse_(param.Search, param.Page); err == nil {
			maxPage := 0
			if len(records) > 0 {
				maxPage = records[0].MaxPage
			}

			helper.WriteResponse(w, http.StatusOK, map[string]interface{}{"max_page": maxPage, "records": records}, nil)
		} else {
			helper.WriteResponse(w, http.StatusInternalServerError, nil, model.ErrInternalServerError)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}
}

func (cinema Cinema) MovieGet(w http.ResponseWriter, r *http.Request) {
	if authorization := cinema.Authorization.Authorization(r); authorization == nil {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	if id, err := strconv.Atoi(mux.Vars(r)["id"]); err == nil {
		if record, err := cinema.MovieGet_(id); err == nil {
			helper.WriteResponse(w, http.StatusOK, record, nil)
		} else {
			helper.WriteResponse(w, http.StatusNotFound, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, model.ErrBadParamInput)
	}
}

func (cinema Cinema) MovieUpdate(w http.ResponseWriter, r *http.Request) {
	if authorization := cinema.Authorization.Authorization(r); authorization == nil || !authorization.IsAdmin {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	record := model.Movie{}
	if err := decoder.Decode(&record); err == nil {
		if err := cinema.MovieUpdate_(&record); err == nil {
			helper.WriteResponse(w, http.StatusOK, record, nil)
		} else {
			helper.WriteResponse(w, http.StatusNotFound, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}
}

func (cinema Cinema) MovieDelete(w http.ResponseWriter, r *http.Request) {
	if authorization := cinema.Authorization.Authorization(r); authorization == nil || !authorization.IsAdmin {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	if id, err := strconv.Atoi(mux.Vars(r)["id"]); err == nil {
		if err := cinema.MovieDelete_(id); err == nil {
			helper.WriteResponse(w, http.StatusOK, nil, nil)
		} else {
			helper.WriteResponse(w, http.StatusNotFound, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, model.ErrBadParamInput)
	}
}

func (cinema Cinema) MovieGetByTitle(w http.ResponseWriter, r *http.Request) {
	if authorization := cinema.Authorization.Authorization(r); authorization == nil {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	title := mux.Vars(r)["title"]

	if record, err := cinema.MovieGetByTitle_(title); err == nil {
		helper.WriteResponse(w, http.StatusOK, record, nil)
	} else {
		helper.WriteResponse(w, http.StatusNotFound, nil, err)
	}
}

func (cinema Cinema) MostViewedMovie(w http.ResponseWriter, r *http.Request) {
	if authorization := cinema.Authorization.Authorization(r); authorization == nil {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	if record, err := cinema.MostViewedMovie_(); err == nil {
		helper.WriteResponse(w, http.StatusOK, record, nil)
	} else {
		helper.WriteResponse(w, http.StatusNotFound, nil, err)
	}
}

func (cinema Cinema) MostViewedGendre(w http.ResponseWriter, r *http.Request) {
	if authorization := cinema.Authorization.Authorization(r); authorization == nil {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	if record, err := cinema.MostViewedGendre_(); err == nil {
		helper.WriteResponse(w, http.StatusOK, record, nil)
	} else {
		helper.WriteResponse(w, http.StatusNotFound, nil, err)
	}
}

func (cinema Cinema) MovieWatch(w http.ResponseWriter, r *http.Request) {
	authorization := cinema.Authorization.Authorization(r)
	if authorization == nil {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	if id, err := strconv.Atoi(mux.Vars(r)["id"]); err == nil {
		if err := cinema.MovieWatch_(id, authorization.UserID); err == nil {
			helper.WriteResponse(w, http.StatusOK, nil, nil)
		} else {
			helper.WriteResponse(w, http.StatusInternalServerError, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, model.ErrBadParamInput)
	}
}

func (cinema Cinema) MovieVote(w http.ResponseWriter, r *http.Request) {
	authorization := cinema.Authorization.Authorization(r)
	if authorization == nil {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	record := model.RatingDetail{}

	if err := decoder.Decode(&record); err == nil {
		record.UserID = authorization.UserID
		if err := cinema.MovieVote_(record); err == nil {
			helper.WriteResponse(w, http.StatusOK, nil, nil)
		} else {
			helper.WriteResponse(w, http.StatusBadRequest, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}
}

func (cinema Cinema) MovieUnvote(w http.ResponseWriter, r *http.Request) {
	authorization := cinema.Authorization.Authorization(r)
	if authorization == nil {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	decoder := json.NewDecoder(r.Body)
	record := model.RatingDetail{}
	if err := decoder.Decode(&record); err == nil {
		if err := cinema.MovieUnvote_(record.MovieID, authorization.UserID); err == nil {
			helper.WriteResponse(w, http.StatusOK, nil, nil)
		} else {
			helper.WriteResponse(w, http.StatusBadRequest, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}
}

func (cinema Cinema) MovieListVoted(w http.ResponseWriter, r *http.Request) {
	authorization := cinema.Authorization.Authorization(r)
	if authorization == nil {
		cinema.Authorization.WriteUnauthorized(w)
		return
	}

	if records, err := cinema.MovieListVoted_(authorization.UserID); err == nil {
		helper.WriteResponse(w, http.StatusOK, map[string]interface{}{"records": records}, nil)
	} else {
		helper.WriteResponse(w, http.StatusInternalServerError, nil, err)
	}
}
