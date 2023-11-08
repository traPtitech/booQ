package model

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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

const (
	BorrowItem = iota
	ReturnItem
	AddItem
	ReduceItem
)

// 0: 借りる
//         1: 返却
//         2: 追加
//         3: 減らす

type RequestPostLogsBody struct {
	OwnerID uint   `json:"ownerId"`
	Type    int    `json:"type"`
	Purpose string `json:"purpose,omitempty"`
	DueDate string `json:"dueDate"`
	Count   int    `json:"count"`
}

// TableName dbのテーブル名を指定する
func (log *Log) TableName() string {
	return "logs"
}

func checkLogsType(value interface{}) error {
	t, _ := value.(int)
	if !(t == 0 || t == 1 || t == 2 || t == 3) {
		return errors.New("must be 1, 2, or 3")
	}
	return nil
}

func (body RequestPostLogsBody) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(&body.OwnerID, validation.Required),
		validation.Field(&body.Type, validation.By(checkLogsType)),
		validation.Field(&body.DueDate, validation.Date("2006-01-02")),
		validation.Field(&body.Count, validation.Required),
	)
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
