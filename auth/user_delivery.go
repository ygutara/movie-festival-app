package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ygutara/movie-festival-app/helper"
)

func (auth Auth) MovieRoute(r *mux.Router) {
	r.HandleFunc("/user/register", auth.Register).Methods("POST")
	r.HandleFunc("/user/login", auth.Login).Methods("POST")
	r.HandleFunc("/user/logout", auth.Logout).Methods("POST")
}

func (auth Auth) Register(w http.ResponseWriter, r *http.Request) {
	if authorization := auth.Authorization(r); authorization != nil {
		helper.WriteResponse(w, http.StatusBadRequest, nil, errors.New("log in already."))
		return
	}

	decoder := json.NewDecoder(r.Body)
	user := UserRegisterOrLogin{}

	if err := decoder.Decode(&user); err == nil {
		// DEBT token cookie
		if _, err := auth.Register_(user); err == nil {
			helper.WriteResponse(w, http.StatusOK, nil, nil)
		} else {
			helper.WriteResponse(w, http.StatusInternalServerError, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}
}

func (auth Auth) Login(w http.ResponseWriter, r *http.Request) {
	if authorization := auth.Authorization(r); authorization != nil {
		helper.WriteResponse(w, http.StatusBadRequest, nil, errors.New("log in already."))
		return
	}

	decoder := json.NewDecoder(r.Body)
	user := UserRegisterOrLogin{}

	if err := decoder.Decode(&user); err == nil {
		// DEBT token cookie
		if _, err := auth.Login_(user); err == nil {
			helper.WriteResponse(w, http.StatusOK, nil, nil)
		} else {
			helper.WriteResponse(w, http.StatusInternalServerError, nil, err)
		}
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}
}

func (auth Auth) Logout(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")

	token := auth.GetBearerToken(reqToken)

	if err := auth.Logout_(token); err == nil {
		helper.WriteResponse(w, http.StatusOK, nil, nil)
	} else {
		helper.WriteResponse(w, http.StatusBadRequest, nil, err)
	}

}
