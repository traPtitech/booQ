package model

import (
	 "errors"

	"github.com/jinzhu/gorm"
)

// Comment commentの構造体
type Comment struct {
	gorm.Model
	ItemID uint `gorm:"type:int;not null" json:"item_id"`
	UserID uint    `gorm:"type:int;not null" json:"user_id"`
	Text   string `gorm:"type:text;not null" json:"text"`
}
type RequestPostCommentBody struct {
	Text   string `gorm:"type:text;not null" json:"text"`
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
	db.Create(&comment)
	return comment, nil
}
