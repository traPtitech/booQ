package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Item itemの構造体
type Item struct {
	gorm.Model
	Name        string `gorm:"type:varchar(64);not null" json:"name"`
	Type        int    `gorm:"type:int;not null" json:"type"`
	Code        string `gorm:"type:varchar(13);" json:"code"`
	Description string `gorm:"type:text;" json:"description"`
	ImgURL      string `gorm:"type:text;" json:"img_url"`
	Owners      []User `gorm:"many2many:item_owners;"`
	Tags        []Tag  `gorm:"many2many:item_tags;"`
	LikeUsers   []User `gorm:"many2many:item_like_users;"`
}

type RequestPostOwnersBody struct {
	UserID     int  `json:"user_id"`
	Rentalable bool `json:"rentalable"`
}

// TableName dbのテーブル名を指定する
func (item *Item) TableName() string {
	return "items"
}

// GetItemByID IDからitemを取得する
func GetItemByID(id int) (Item, error) {
	res := Item{}
	db.Where("id = ?", id).First(&res)
	if res.Name == "" {
		return Item{}, errors.New("Nameが不正です")
	}
	return res, nil
}

// GetItems 全itemを取得する
func GetItems() ([]Item, error) {
	res := []Item{}
	db.Find(&res)
	return res, nil
}

// CreateItem 新しいItemを登録する
func CreateItem(item Item) (Item, error) {
	if item.Name == "" {
		return Item{}, errors.New("Nameが存在しません")
	}
	db.Create(&item)
	return item, nil
}

// RegisterItem 新しい所有者を登録する
func RegisterOwner(user User, item Item) (Item, error) {
	if user.Name == "" {
		return Item{}, errors.New("UserのNameが存在しません")
	}
	if item.Name == "" {
		return Item{}, errors.New("ItemのNameが存在しません")
	}
	db.Model(&item).Association("Owners").Append(&user)
	return item, nil
}
