package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
)

// User userの構造体
type User struct {
	GormModel
	Name        string `gorm:"type:varchar(32);unique;not null;size:50" json:"name"`
	DisplayName string `gorm:"type:varchar(64);not null" json:"displayName"`
	Admin       bool   `gorm:"default:false" json:"admin"`
}

type RequestPutUsersBody struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

// TableName dbのテーブル名を指定する
func (user *User) TableName() string {
	return "users"
}

func (body RequestPutUsersBody) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(&body.Name, validation.Required),
		validation.Field(&body.Admin, validation.Skip),
	)
}

// GetUsers 全userを取得する
func GetUsers() ([]User, error) {
	res := []User{}
	err := db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserByName userをNameから取得する
func GetUserByName(name string) (User, error) {
	res := User{}
	err := db.Where("name = ?", name).First(&res).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, errors.New("nameが存在しません")
		}
		return User{}, err
	}
	return res, nil
}

// GetUserByID userをIDから取得する
func GetUserByID(id int) (User, error) {
	res := User{}
	err := db.Where("id = ?", id).First(&res).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, errors.New("該当するNameがありません")
		}
		return User{}, err
	}
	return res, nil
}

// CreateUser userを作成する
func CreateUser(user User) (User, error) {
	if user.Name == "" {
		return User{}, errors.New("Nameが存在しません")
	}
	err := db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// UpdateUser userの情報(表示される名前やアイコン、管理者権限)の変更
func UpdateUser(newUser User) (User, error) {
	user := User{}
	if newUser.Name == "" {
		return User{}, errors.New("Nameが存在しません")
	}
	err := db.Where("name = ?", newUser.Name).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{}, errors.New("指定したuserが存在しません")
		}
		return User{}, err
	}
	user.DisplayName = newUser.DisplayName
	user.Admin = newUser.Admin
	err = db.Save(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// CheckTargetedOrAdmin 管理者または対象のUserかどうかを判定する
func CheckTargetedOrAdmin(user, reqUser User) error {
	if user.Name == "" || reqUser.Name == "" {
		return errors.New("Nameが存在しません")
	}
	if !user.Admin && reqUser.Name != user.Name {
		return errors.New("あなたは管理者でもなければ対象のUser本人でもありません")
	}
	return nil
}

// CheckOwnsOrAdmin 管理者またはItemの所有者かどうかを判定する
func CheckOwnsOrAdmin(user *User, item *Item) error {
	if user.Admin {
		return nil
	}
	for _, owner := range item.Owners {
		if owner.UserID == user.ID {
			return nil
		}
	}
	return errors.New("あなたは管理者でもなければ対象のItemの所有者でもありません")
}
