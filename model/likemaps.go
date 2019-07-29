package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Likemaps likemapsの構造体
type Likemaps struct {
	gorm.Model
	ItemID      int		`json:"item_id"`
	UserID		int		`json:"user_id"`
}

// // TableName dbのテーブル名を指定する
// func (likemaps *Likemaps) TableName() string {
// 	return "likemapss"
// }