package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-"`
	IsAdmin  bool   `json:"-"`
}

func (User) TableName() string {
	return "user"
}

type UserToken struct {
	ID     int `gorm:"primaryKey"`
	UserID int
	Token  string
}

func (UserToken) TableName() string {
	return "user_token"
}
