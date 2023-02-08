package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/ygutara/movie-festival-app/model"
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

type UserRegisterOrLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthorizationData struct {
	UserID  int
	IsAdmin bool
}

func (auth Auth) Authorization(r *http.Request) (authorization *AuthorizationData) {
	reqToken := r.Header.Get("Authorization")

	token := auth.GetBearerToken(reqToken)

	userToken, _ := auth.UserTokenGet(token)
	if userToken.User != nil {
		authorization.UserID = userToken.User.ID
		authorization.IsAdmin = userToken.User.IsAdmin
	}

	return
}

func (auth Auth) WriteUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	response, _ := json.Marshal(map[string]string{"error_desciption": " You are not authorized"})
	w.Write(response)
}

func (auth Auth) GetBearerToken(reqToken string) (token string) {
	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) >= 2 {
		token = splitToken[1]
	}

	return
}

func (auth Auth) UserTokenGet(token string) (userToken model.UserToken, err error) {
	db := auth.DB

	if result := db.Preload("User").Where("token = ? AND is_active = ?", token, true).First(&userToken); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("token not applicable")
	}

	return
}
