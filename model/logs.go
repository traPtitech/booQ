package model

import (
	// "errors"
	"github.com/jinzhu/gorm"
	"time"
)

// Log logの構造体
type Log struct {
	gorm.Model
	ItemID  int       `cjson:"item_id"`
	UserID  int       `gorm:"type:int;not null" json:"user_id"`
	OwnerID int       `gorm:"type:int;not null" json:"owner_id"`
	Type    int       `gorm:"type:int;not null" json:"type"`
	Purpose time.Time `json:"purpose"`
	DueDate string    `gorm:"type:datetime;" json:"due_date"`
}

// TableName dbのテーブル名を指定する
func (log *Log) TableName() string {
	return "logs"
}
