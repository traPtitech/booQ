package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Comment commentの構造体
type Comment struct {
	GormModel
	ItemID uint   `gorm:"type:int;not null" json:"itemId"`
	UserID uint   `gorm:"type:int;not null" json:"userId"`
	User   User   `json:"user"`
	Text   string `gorm:"type:text;not null" json:"text"`
}
type ResponseGetComments struct {
	*Comment
	Item Item `json:"item"`
}
type RequestPostCommentBody struct {
	Text string `gorm:"type:text;not null" json:"text"`
}

func (body RequestPostCommentBody) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(&body.Text, validation.Required),
	)
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
	err = db.Preload("User").Create(&comment).Error
	if err != nil {
		return Comment{}, err
	}
	return comment, nil
}

// GetCommentsByUserID UserIDからCommentsを取得する
func GetCommentsByUserID(userID uint) ([]ResponseGetComments, error) {
	comments := []Comment{}
	err := db.Preload("User").Find(&comments, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	ids := []uint{}
	for _, c := range comments {
		ids = append(ids, c.ItemID)
	}
	items := []Item{}
	res := []ResponseGetComments{}
	err = db.Where(ids).Find(&items).Error
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		for _, comment := range comments {
			if item.ID == comment.ItemID {
				res = append(res, ResponseGetComments{&comment, item})
				break
			}
		}
	}
	return res, nil
}
