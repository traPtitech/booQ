package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Comment commentの構造体
type Comment struct {
	gorm.Model
	ItemID string `gorm:"type:int;not null" json:"item_id"`
	UserID int    `gorm:"type:int;not null" json:"user_id"`
	Text   string `gorm:"type:text;not null" json:"text"`
}

// TableName dbのテーブル名を指定する
func (comment *Comment) TableName() string {
	return "comments"
}
