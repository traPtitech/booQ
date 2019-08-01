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
	Admin       bool   `gorm:"default:false" json:"admin"`
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

// UpdateUser userに管理者権限を付与する
func UpdateUser(user User,newUser User) (User, error) {
	nowUser := User{}
	db.Where("name = ?", newUser.Name).First(&nowUser)
	if nowUser.IconFileID == newUser.IconFileID && nowUser.DisplayName == newUser.DisplayName && nowUser.Admin == newUser.Admin  {
		return User{}, errors.New("更新されるべき情報がありません")
	}
	res := User{}
	db.Model(&res).Where("name = ?", newUser.Name).Updates(newUser)
	return res, nil
}
