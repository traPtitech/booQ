package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Tag tagの構造体
type Tag struct {
	gorm.Model
	Name      string	`json:"name"`
}

// // TableName dbのテーブル名を指定する
// func (tag *Tag) TableName() string {
// 	return "tags"
// }