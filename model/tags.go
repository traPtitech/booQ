package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Tag tagの構造体
type Tag struct {
	gorm.Model
	Name  string `gorm:"type:varchar(32);unique;not null" json:"name"`
	Items []Item `gorm:"many2many:items_tags;" json:"items"`
}

type RequestPostTagsBody struct {
	Name string `json:"name"`
}

type RequestPostTags2ItemBody struct {
	ID []uint `json:"id"`
}

// TableName dbのテーブル名を指定する
func (tag *Tag) TableName() string {
	return "tags"
}

// CreateTag Tagを作成する
func CreateTag(name string) (Tag, error) {
	if name == "" {
		return Tag{}, errors.New("nameが存在しません")
	}
	tag := Tag{Name: name}
	db.Create(&tag)
	return tag, nil
}

// AttachTag Itemにタグをつける
func AttachTag(tagID, itemID uint) (Tag, error) {
	tag := Tag{}
	item := Item{}
	db.First(&tag, tagID)
	db.First(&item, itemID)
	if tag.ID == 0 || item.ID == 0 {
		return Tag{}, errors.New("指定されたタグ、またはアイテムは存在しません")
	}
	db.Model(&tag).Association("Items").Append(&item)
	return tag, nil
}

// RemoveTag ItemからTagを外す
func RemoveTag(tag Tag, itemID uint) (Tag, error) {
	if tag.ID == 0 || itemID == 0 {
		return Tag{}, errors.New("タグまたはアイテムが指定されていません")
	}
	existed := false
	for _, tagItem := range tag.Items {
		if tagItem.ID == itemID {
			existed = true
		}
	}
	if !existed {
		return Tag{}, errors.New("指定のアイテムに指定のタグはついていません")
	}
	item := Item{}
	item.ID = itemID
	db.Model(&tag).Association("Items").Delete(&item)
	return tag, nil
}

// GetTags すべてのタグを取得する(Itemについての情報は無し)
func GetTags() ([]Tag, error) {
	tags := []Tag{}
	db.Find(&tags)
	return tags, nil
}

// GetTagByName 一つのタグをNameから取得する(Itemについての詳しい情報は無し)
func GetTagByName(name string) (Tag, error) {
	tag := Tag{}
	db.Find(&tag, "name = ?", name).Related(&tag.Items, "Items")
	if tag.ID == 0 {
		return Tag{}, errors.New("指定されたタグ、またはアイテムは存在しません")
	}
	return tag, nil
}

// GetTagByID 一つのタグをIDから取得する(Itemについての詳しい情報は無し)
func GetTagByID(tagID uint) (Tag, error) {
	tag := Tag{}
	db.Find(&tag, "id = ?", tagID).Related(&tag.Items, "Items")
	if tag.ID == 0 {
		return Tag{}, errors.New("指定されたタグ、またはアイテムは存在しません")
	}
	return tag, nil
}
