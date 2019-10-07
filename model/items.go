package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Item itemの構造体
type Item struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(64);not null" json:"name"`
	Type        int       `gorm:"type:int;not null" json:"type"`
	Code        string    `gorm:"type:varchar(13);" json:"code"`
	Description string    `gorm:"type:text;" json:"description"`
	ImgURL      string    `gorm:"type:text;" json:"img_url"`
	Owners      []Owner   `gorm:"many2many:ownership_maps;" json:"owners"`
	Logs        []Log     `json:"logs"`
	LatestLogs  []Log     `json:"latest_logs"`
	Comments    []Comment `json:"commments"`
	Likes       []User    `gorm:"many2many:like_maps;" json:"likes"`
}

type Owner struct {
	gorm.Model
	UserID     uint `gorm:"type:int;not null" json:"owner_id"`
	User       User `json:"owner"`
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
	res.ID = id
	db.Set("gorm:auto_preload", true).First(&res).Related(&res.Owners, "Owners").Related(&res.Logs, "Logs").Related(&res.Comments, "Comments").Related(&res.Likes, "Likes")
	if res.Name == "" {
		return Item{}, errors.New("該当するItemがありません")
	}
	var err error
	res.LatestLogs, err = GetLatestLogs(id)
	if err != nil {
		return Item{}, err
	}
	return res, nil
}

// GetItemByName Nameからitemを取得する
func GetItemByName(name string) (Item, error) {
	res := Item{}
	db.Set("gorm:auto_preload", true).First(&res, "name = ?", name).Related(&res.Owners, "Owners").Related(&res.Logs, "Logs").Related(&res.Comments, "Comments").Related(&res.Likes, "Likes")
	if res.Name == "" {
		return Item{}, errors.New("該当するNameがありません")
	}
	var err error
	res.LatestLogs, err = GetLatestLogs(res.ID)
	if err != nil {
		return Item{}, err
	}
	return res, nil
}

// GetItems 全itemを取得する
func GetItems() ([]Item, error) {
	res := []Item{}
	db.Find(&res)
	for i, item := range res {
		db.Set("gorm:auto_preload", true).First(&item).Related(&item.Owners, "Owners").Related(&item.Logs, "Logs").Related(&item.Comments, "Comments").Related(&item.Likes, "Likes")
		var err error
		item.LatestLogs, err = GetLatestLogs(item.ID)
		if err != nil {
			return []Item{}, err
		}
		res[i] = item
	}
	return res, nil
}

// CreateItem 新しいItemを登録する
func CreateItem(item Item) (Item, error) {
	if item.Name == "" {
		return Item{}, errors.New("Nameが存在しません")
	}
	reddiedItem := Item{}
	db.Where("name = ?", item.Name).Or("code != '' AND code = ?", item.Code).Find(&reddiedItem)

	if reddiedItem.Name != "" {
		return Item{}, errors.New("すでに同じItemが存在しています")
	}
	db.Create(&item)
	return item, nil
}

// RegisterOwner 新しい所有者を登録する
func RegisterOwner(owner Owner, item Item) (Item, error) {
	var existed bool
	db.Preload("Owners").Find(&item)
	owner.User, _ = GetUserByID(int(owner.UserID))
	for _, nowOwner := range item.Owners {
		if nowOwner.UserID != owner.UserID {
			continue
		}
		if owner.Rentalable == nowOwner.Rentalable {
			nowOwner.Count += owner.Count
		} else {
			nowOwner.Count = owner.Count
		}
		existed = true
		nowOwner.User = owner.User
		db.Model(&item).Association("Owners").Replace(&nowOwner)
	}
	if !existed {
		db.Create(&owner)
		db.Model(&item).Association("Owners").Append(&owner)
	}
	return item, nil
}

// PushLike likeを押す
func PushLike(itemID, userID uint) (Item, error) {
	var existed bool
	item := Item{}
	db.Set("gorm:auto_preload", true).First(&item, itemID).Related(&item.Likes, "Likes")
	user, _ := GetUserByID(int(userID))
	for _, likeUser := range item.Likes {
		if likeUser.ID == userID {
			existed = true

		}
	}
	if existed {
		return Item{}, errors.New("すでにいいねしています")
	} else {
		db.Model(&item).Association("Likes").Append(&user)
	}
	return item, nil
}
