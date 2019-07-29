package model

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

// Log logの構造体
type Log struct {
	gorm.Model
	ItemID      int		`json:"item_id"`
	UserID		int		`json:"user_id"`
	OwnerID		int		`json:"owner_id"`
	Type		int		`json:"type"`
	Purpose		string	`json:"purpose"`
	DueDate		string	`json:"due_date"`
}

// // TableName dbのテーブル名を指定する
// func (log *Log) TableName() string {
// 	return "logs"
// }