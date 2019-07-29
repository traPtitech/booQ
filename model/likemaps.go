package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Likemaps likemapsの構造体
type Likemap struct {
	gorm.Model
	ItemID int `gorm:"type:int;not null" json:"item_id"`
	UserID int `gorm:"type:int;not null" json:"user_id"`
}

// TableName dbのテーブル名を指定する
func (likemap *Likemap) TableName() string {
	return "likemaps"
}
