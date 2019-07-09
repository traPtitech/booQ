package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User userの構造体
type User struct {
	gorm.Model
	Name       string    `gorm:"unique;not null;size:50" json:"name"`
	ScreenName string    `json:"screen_name"`
	IconFileID string    `json:"icon_file_id"`
	Admin      int       `gorm:"default:0" json:"admin"`
}

// GetUser userを取得する
func GetUser(user User) (User, error) {
	if user.Name == "" {
		return User{}, errors.New("Nameが存在しません")
	}
	res := User{}
	db.Where("name = ?", user.Name).First(&res)
	return res, nil
}

// CreateUser userを作成する
func CreateUser(user User) (User, error) {
	if user.Name == "" {
		return User{}, errors.New("Nameが存在しません")
	}
	db.Create(&user)
	return user, nil
}
