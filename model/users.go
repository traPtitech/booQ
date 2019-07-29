package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User userの構造体
type User struct {
	gorm.Model
	Name        string `gorm:"unique;not null;size:50" json:"name"`
	DisplayName string `json:"displayName"`
	IconFileID  string `json:"iconFileId"`
	Admin       bool    `gorm:"default:false" json:"admin"`
}

// TableName dbのテーブル名を指定する
func (user *User) TableName() string {
	return "users"
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
