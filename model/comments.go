package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Comment commentの構造体
type Comment struct {
	gorm.Model
	ItemID      string	`json:"item_id"`
	UserID		int		`json:"user_id"`
	Text		string	`json:"text"`
}

// // TableName dbのテーブル名を指定する
// func (comment *Comment) TableName() string {
// 	return "comments"
// }