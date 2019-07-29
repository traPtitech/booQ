package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Item itemの構造体
type Item struct {
	gorm.Model
	Name        string	`json:"name"`
	Type		int		`json:"type"`
	Code		int		`json:"code"`
	Description	string	`json:"description"`
	ImgUrl		string	`json:"img_url"`

}

// // TableName dbのテーブル名を指定する
// func (item *Item) TableName() string {
// 	return "items"
// }

// // GetItems 全itemを取得する
// func GetItems() ([]Item, error) {
// }

