package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Log logの構造体
type Log struct {
	gorm.Model
	ItemID  int       `gorm:"type:int;not null" json:"item_id"`
	UserID  int       `gorm:"type:int;not null" json:"user_id"`
	OwnerID int       `gorm:"type:int;not null" json:"owner_id"`
	Type    int       `gorm:"type:int;not null" json:"type"`
	Purpose string    `json:"purpose"`
	DueDate time.Time `gorm:"type:datetime;" json:"due_date"`
	Count   int       `gorm:"type:int;not null" json:"count"`
}

type RequestPostLogsBody struct {
	OwnerID int       `json:"owner_id"`
	Type    int       `json:"type"`
	Purpose string    `json:"purpose"`
	DueDate time.Time `json:"due_date"`
	Count   int       `json:"count"`
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

// GetLatestLog OwnerIDからLatestLogを取得する
func GetLatestLog(itemID, ownerID int) (Log, error) {
	item, err := GetItemByID(itemID)
	if err != nil {
		return Log{}, err
	}
	exist := false
	for _, owner := range item.Owners {
		if int(owner.ID) == ownerID {
			exist = true
		}
	}
	if !exist {
		return Log{}, errors.New("指定した所有者はそのItemを所有していません")
	}
	log := Log{}
	db.Last(&log).Where("item_id = ? AND owner_id = ?", itemID, ownerID)
	return log, nil
}
