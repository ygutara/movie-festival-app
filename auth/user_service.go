package auth

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/ygutara/movie-festival-app/model"
	"gorm.io/gorm"
)

func (auth Auth) Register_(userTemplate UserRegisterOrLogin) (token string, err error) {
	userExisted, _ := auth.UserGet(userTemplate.Username)
	if userExisted.Username == "" {
		err = errors.New("username can not be blank")
		return
	} else if userExisted.Username == "" {
		err = errors.New("username has been used")
		return
	}

	user := model.User{}
	user.Username = userTemplate.Username
	user.Password = userTemplate.Password

	auth.DB.Create(&user)

	auth.UserTokenCreate(user.ID)

	return
}

func (auth Auth) Login_(userTemplate UserRegisterOrLogin) (token string, err error) {
	userExisted, _ := auth.UserGet(userTemplate.Username)

	if userExisted.Password != userTemplate.Password {
		err = errors.New("username or password wrong")
		return
	}

	auth.UserTokenCreate(userExisted.ID)

	return
}

func (auth Auth) Logout_(token string) (err error) {
	token = auth.GetBearerToken(token)
	db := auth.DB
	userToken := model.UserToken{}

	userToken, err = auth.UserTokenGet(token)
	if err != nil {
		return
	}

	db.Model(&model.UserToken{}).Where("token = ?", userToken.Token).Update("is_active", false)

	return
}

func (auth Auth) UserGet(username string) (user model.User, err error) {
	auth.DB.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		err = model.ErrNotFound
	}

	return
}

func (auth Auth) UserTokenCreate(userId int) (token string) {
	db := auth.DB
	different := false
	for !different {
		if result := db.Where("token = ? AND is_active = ?", token, true).First(&model.UserToken{}); !errors.Is(result.Error, gorm.ErrRecordNotFound) || token == "" {
			token, _ = GenerateToken()
		} else {
			different = true
		}
	}

	userToken := model.UserToken{}
	userToken.UserID = userId
	userToken.IsActive = true
	userToken.Token = token

	auth.DB.Create(&userToken)

	return
}

func GenerateToken() (string, bool) {
	return uuid.NewV4().String(), false
}
