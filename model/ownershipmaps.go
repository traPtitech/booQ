package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Ownershipmaps ownershipmapsの構造体
type Ownershipmaps struct {
	gorm.Model
	ItemID      int		`json:"item_id"`
	UserID		int		`json:"user_id"`
	Rentalable	bool	`json:"rentalable"`
}

// // TableName dbのテーブル名を指定する
// func (ownershipmaps *Ownershipmaps) TableName() string {
// 	return "ownershipmapss"
// }