package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Log logの構造体
type Log struct {
	gorm.Model
	ItemID  uint      `gorm:"type:int;not null" json:"item_id"`
	UserID  uint      `gorm:"type:int;not null" json:"user_id"`
	User    User      `json:"user"`
	OwnerID uint      `gorm:"type:int;not null" json:"owner_id"`
	Owner   User      `json:"owner"`
	Type    int       `gorm:"type:int;not null" json:"type"`
	Purpose string    `json:"purpose"`
	DueDate time.Time `gorm:"type:datetime;" json:"due_date"`
	Count   int       `gorm:"type:int;not null" json:"count"`
}

type RequestPostLogsBody struct {
	OwnerID uint   `json:"owner_id"`
	Type    int    `json:"type"`
	Purpose string `json:"purpose"`
	DueDate string `json:"due_date"`
	Count   int    `json:"count"`
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
	_, err = GetUserByID(int(log.OwnerID))
	if err != nil {
		return Log{}, errors.New("Ownerが存在しません")
	}
	db.Create(&log)
	return log, nil
}

// GetLatestLog ownerIDからLatestLogを取得する
func GetLatestLog(logs []Log, ownerID uint) (Log, error) {
	logMap := map[uint]Log{}
	for _, log := range logs {
		nowLatestLog, exist := logMap[log.OwnerID]
		if exist {
			if !nowLatestLog.CreatedAt.After(log.CreatedAt) {
				logMap[log.OwnerID] = log
			}
		} else {
			logMap[log.OwnerID] = log
		}
	}
	latestLog := logMap[ownerID]
	return latestLog, nil
}

// GetLatestLogs 各所有者ごとの最新のログを取得する。
func GetLatestLogs(logs []Log) ([]Log, error) {
	logMap := map[uint]Log{}
	for _, log := range logs {
		nowLatestLog, exist := logMap[log.OwnerID]
		if exist {
			if !nowLatestLog.CreatedAt.After(log.CreatedAt) {
				logMap[log.OwnerID] = log
			}
		} else {
			logMap[log.OwnerID] = log
		}
	}
	latestLogs := []Log{}
	for _, latestLog := range logMap {
		latestLogs = append(latestLogs, latestLog)
	}
	return latestLogs, nil
}

// GetLogsByItemID itemIDからLogsを取得する
func GetLogsByItemID(itemID uint) ([]Log, error) {
	// 指定のitemIDのitemが存在するかどうかはここで判別つけていません
	logs := []Log{}
	db.Set("gorm:auto_preload", true).Find(&logs, "item_id = ?", itemID)
	return logs, nil
}
