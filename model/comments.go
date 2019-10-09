package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Comment commentの構造体
type Comment struct {
	gorm.Model
	ItemID uint   `gorm:"type:int;not null" json:"item_id"`
	UserID uint   `gorm:"type:int;not null" json:"user_id"`
	User   User   `json:"user"`
	Text   string `gorm:"type:text;not null" json:"text"`
}
type RequestPostCommentBody struct {
	Text string `gorm:"type:text;not null" json:"text"`
}

// TableName dbのテーブル名を指定する
func (comment *Comment) TableName() string {
	return "comments"
}

// CreateComments 新しいCommentを登録する
func CreateComment(comment Comment) (Comment, error) {
	if comment.ItemID == 0 {
		return Comment{}, errors.New("ItemIDが存在しません")
	}
	_, err := GetItemByID(comment.ItemID)
	if err != nil {
		return Comment{}, errors.New("Itemが存在しません")
	}

	if comment.Text == "" {
		return Comment{}, errors.New("Textが存在しません")
	}
	db.Preload("User").Create(&comment)
	return comment, nil
}

// GetCommentsByUserID UserIDからCommentsを取得する
func GetCommentsByUserID(userID uint) ([]Comment, error) {
	comments := []Comment{}
	db.Preload("User").Find(&comments, "user_id = ?", userID)
	return comments, nil
}
