package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Tagmap tagmapの構造体
type Tagmap struct {
	gorm.Model
	ItemID int `gorm:"type:int;not null" json:"item_id"`
	TagID  int `gorm:"type:int;not null" json:"tag_id"`
}

// TableName dbのテーブル名を指定する
func (tagmap *Tagmap) TableName() string {
	return "tag_maps"
}
