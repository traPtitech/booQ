package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Ownershipmaps ownershipmapsの構造体
type Ownershipmap struct {
	gorm.Model
	ItemID     int  `gorm:"type:int;not null" json:"item_id"`
	UserID     int  `gorm:"type:int;not null" json:"user_id"`
	Rentalable bool `gorm:"default:true" json:"rentalable"`
}

// TableName dbのテーブル名を指定する
func (ownershipmap *Ownershipmap) TableName() string {
	return "ownershipmaps"
}
