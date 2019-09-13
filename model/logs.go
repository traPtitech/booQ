package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Log logの構造体
type Log struct {
	gorm.Model
	ItemID  int       `json:"item_id"`
	UserID  int       `gorm:"type:int;not null" json:"user_id"`
	OwnerID int       `gorm:"type:int;not null" json:"owner_id"`
	Type    int       `gorm:"type:int;not null" json:"type"`
	Purpose string    `json:"purpose"`
	DueDate time.Time `gorm:"type:datetime;" json:"due_date"`
}

type RequestPostLogsBody struct {
	OwnerID int       `json:"owner_id"`
	Type    int       `json:"type"`
	Purpose string    `json:"purpose"`
	DueDate time.Time `json:"due_date"`
}

// TableName dbのテーブル名を指定する
func (log *Log) TableName() string {
	return "logs"
}

// CreateLog logを作成する
func CreateLog(log Log) (Log, error) {
	if log.ItemID == 0 {
		return Log{}, errors.New("ItemIDが存在しません")
	}
	_, err := GetItemByID(log.ItemID)
	if err != nil {
		return Log{}, errors.New("Itemが存在しません")
	}
	_, err = GetUserByID(log.OwnerID)
	if err != nil {
		return Log{}, errors.New("Ownerが存在しません")
	}
	db.Create(&log)
	return log, nil
}
