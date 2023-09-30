package model

import (
	"errors"
	"time"
)

// Log logの構造体
type Log struct {
	GormModel
	ItemID  uint      `gorm:"type:int;not null" json:"itemId"`
	UserID  uint      `gorm:"type:int;not null" json:"userId"`
	User    User      `json:"user"`
	OwnerID uint      `gorm:"type:int;not null" json:"ownerId"`
	Owner   User      `json:"owner"`
	Type    int       `gorm:"type:int;not null" json:"type"`
	Purpose string    `json:"purpose"`
	DueDate time.Time `gorm:"type:datetime;" json:"dueDate"`
	Count   int       `gorm:"type:int;not null" json:"count"`
}

type RequestPostLogsBody struct {
	OwnerID uint   `json:"ownerId"`
	Type    int    `json:"type"`
	Purpose string `json:"purpose"`
	DueDate string `json:"dueDate"`
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
	err = db.Create(&log).Error
	if err != nil {
		return Log{}, err
	}
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
