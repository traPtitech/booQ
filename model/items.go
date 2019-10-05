package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Item itemの構造体
type Item struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(64);not null" json:"name"`
	Type        int      `gorm:"type:int;not null" json:"type"`
	Code        string   `gorm:"type:varchar(13);" json:"code"`
	Description string   `gorm:"type:text;" json:"description"`
	ImgURL      string   `gorm:"type:text;" json:"img_url"`
	Owners      []*Owner `gorm:"many2many:ownership_maps;" json:"owners"`
}

type Owner struct {
	gorm.Model
	OwnerID    uint `gorm:"type:int;not null" json:"owner_id"`
	Rentalable bool `gorm:"type:bool;not null" json:"rentalable"`
	Count      int  `gorm:"type:int;default:1" json:"count"`
}

type RequestPostOwnersBody struct {
	UserID     int  `json:"user_id"`
	Rentalable bool `json:"rentalable"`
	Count      int  `json:"count"`
}

// TableName dbのテーブル名を指定する
func (item *Item) TableName() string {
	return "items"
}

func (item *Owner) TableName() string {
	return "owners"
}

// GetItemByID IDからitemを取得する
func GetItemByID(id uint) (Item, error) {
	res := Item{}
	db.First(&res, id).Related(&res.Owners, "Owners")
	if res.Name == "" {
		return Item{}, errors.New("該当するItemがありません")
	}
	return res, nil
}

// GetItemByName Nameからitemを取得する
func GetItemByName(name string) (Item, error) {
	res := Item{}
	db.Where("name = ?", name).First(&res)
	if res.Name == "" {
		return Item{}, errors.New("該当するNameがありません")
	}
	return res, nil
}

// GetItems 全itemを取得する
func GetItems() ([]Item, error) {
	res := []Item{}
	db.Find(&res)
	for i, item := range res {
		itemWithOwner := Item{}
		db.First(&itemWithOwner).Related(&item.Owners, "Owners").Where("name=?", item.Name)
		res[i] = itemWithOwner
	}
	return res, nil
}

// CreateItem 新しいItemを登録する
func CreateItem(item Item) (Item, error) {
	if item.Name == "" {
		return Item{}, errors.New("Nameが存在しません")
	}
	reddiedItem := Item{}
	db.Where("name = ?", item.Name).Or("code = ?", item.Code).Find(&reddiedItem)

	if reddiedItem.Name != "" {
		return Item{}, errors.New("すでに同じItemが存在しています")
	}
	db.Create(&item)
	return item, nil
}

// RegisterItem 新しい所有者を登録する
func RegisterOwner(owner Owner, item Item) (Item, error) {
	var existed bool
	db.Model(&item).Related(&item.Owners, "Owners")
	for _, nowOwner := range item.Owners {
		if nowOwner.OwnerID != owner.OwnerID {
			continue
		}
		if owner.Rentalable == nowOwner.Rentalable {
			nowOwner.Count += owner.Count
		} else {
			nowOwner.Count = owner.Count
		}
		existed = true
		db.Save(&nowOwner)
		db.Model(&item).Related(&item.Owners, "Owners")
	}
	if !existed {
		db.Create(&owner)
		db.Model(&item).Association("Owners").Append(&owner)
	}
	return item, nil
}
