package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Item itemの構造体
type Item struct {
	gorm.Model
	Name        string `gorm:"type:varchar(64);not null" json:"name"`
	Type        int    `gorm:"type:int;not null" json:"type"`
	Code        int    `gorm:"type:int;" json:"code"`
	Description string `gorm:"type:text;" json:"description"`
	ImgURL      string `gorm:"type:text;" json:"img_url"`
}

// TableName dbのテーブル名を指定する
func (item *Item) TableName() string {
	return "items"
}

// // GetItems 全itemを取得する
// func GetItems() ([]Item, error) {
// }
