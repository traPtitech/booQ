package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Item itemの構造体
type Item struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(64);not null" json:"name"`
	Type        int          `gorm:"type:int;not null" json:"type"`
	Code        string       `gorm:"type:varchar(13);" json:"code"`
	Description string       `gorm:"type:text;" json:"description"`
	ImgURL      string       `gorm:"type:text;" json:"img_url"`
	Owners      []Owner      `gorm:"many2many:ownership_maps;" json:"owners"`
	RentalUsers []RentalUser `gorm:"many2many:rental_user_maps;" json:"rental_users"`
	Logs        []Log        `json:"logs"`
	LatestLogs  []Log        `json:"latest_logs"`
	Comments    []Comment    `json:"comments"`
	Likes       []User       `gorm:"many2many:like_maps;" json:"likes"`
	LikeCounts  int          `gorm:"-" json:"like_counts"`
}

type Owner struct {
	gorm.Model
	UserID     uint `gorm:"type:int;not null" json:"owner_id"`
	User       User `json:"user"`
	Rentalable bool `gorm:"type:bool;not null;primary_key" json:"rentalable"`
	Count      int  `gorm:"type:int;default:1" json:"count"`
}

type RentalUser struct {
	gorm.Model
	UserID  uint `gorm:"type:int;not null" json:"user_id"`
	User    User `json:"user"`
	OwnerID uint `gorm:"type:int;not null" json:"owner_id"`
	Owner   User `gorm:"foreignkey:OwnerID" json:"owner"`
	Count   int  `gorm:"type:int;default:1" json:"count"`
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

func (owner *Owner) TableName() string {
	return "owners"
}

func (rentalUser *RentalUser) TableName() string {
	return "rental_users"
}

// GetItemByID IDからitemを取得する
func GetItemByID(id uint) (Item, error) {
	res := Item{}
	db.Set("gorm:auto_preload", true).First(&res, id).Related(&res.Owners, "Owners").Related(&res.Logs, "Logs").Related(&res.RentalUsers, "RentalUsers").Related(&res.Comments, "Comments").Related(&res.Likes, "Likes")
	if res.Name == "" {
		return Item{}, errors.New("該当するItemがありません")
	}
	var err error
	res.LatestLogs, err = GetLatestLogs(res.Logs)
	if err != nil {
		return Item{}, err
	}
	return res, nil
}

// GetItemByName Nameからitemを取得する
func GetItemByName(name string) (Item, error) {
	res := Item{}
	db.Set("gorm:auto_preload", true).First(&res, "name = ?", name).Related(&res.Owners, "Owners").Related(&res.RentalUsers, "RentalUsers").Related(&res.Logs, "Logs").Related(&res.Comments, "Comments").Related(&res.Likes, "Likes")
	if res.Name == "" {
		return Item{}, errors.New("該当するNameがありません")
	}
	var err error
	res.LatestLogs, err = GetLatestLogs(res.Logs)
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
		db.Set("gorm:auto_preload", true).First(&item).Related(&item.Owners, "Owners").Related(&item.Logs, "Logs").Related(&item.RentalUsers, "RentalUsers").Related(&item.Comments, "Comments").Related(&item.Likes, "Likes")
		var err error
		item.LatestLogs, err = GetLatestLogs(item.Logs)
		if err != nil {
			return []Item{}, err
		}
		item.LikeCounts = len(item.Likes)
		item.Likes = []User{}
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
		db.Save(&nowOwner)
		db.Set("gorm:auto_preload", true).First(&item).Related(&item.Owners, "Owners").Related(&item.Logs, "Logs")
		latestLog, err := GetLatestLog(item.Logs, owner.UserID)
		if err != nil {
			return Item{}, err
		}
		if latestLog.ItemID != 0 {
			_, err = CreateLog(Log{ItemID: latestLog.ItemID, UserID: owner.UserID, OwnerID: owner.UserID, Type: 2, Count: latestLog.Count + owner.Count})
			if err != nil {
				return Item{}, err
			}
		}
	}
	if !existed {
		db.Create(&owner)
		db.Model(&item).Association("Owners").Append(&owner)
		_, err := CreateLog(Log{ItemID: item.ID, UserID: owner.UserID, OwnerID: owner.UserID, Type: 2, Count: owner.Count})
		if err != nil {
			return Item{}, err
		}
	}
	return item, nil
}

// RentalItem 物品を借りたりするときにRentalUserを作成する
func RentalItem(rentalUser RentalUser, ownerID uint, item Item, logType int) (Item, error) {
	var existed bool
	db.Preload("RentalUsers").Preload("Owners").Find(&item)
	// owner.User, _ = GetUserByID(int(owner.UserID))
	for _, nowRentalUser := range item.RentalUsers {
		if nowRentalUser.UserID != rentalUser.UserID || nowRentalUser.Owner.ID != ownerID {
			continue
		}
		existed = true
		if logType == 0 {
			nowRentalUser.Count += rentalUser.Count
		}
		if logType == 1 {
			nowRentalUser.Count += rentalUser.Count
		}
		if nowRentalUser.Count > 0 {
			return Item{}, errors.New("Return超過3")
		}
		db.Save(&nowRentalUser)
		db.Set("gorm:auto_preload", true).First(&item).Related(&item.Owners, "Owners").Related(&item.Logs, "Logs").Related(&item.RentalUsers, "RentalUsers")
	}
	if !existed {
		db.Create(&rentalUser)
		db.Model(&item).Association("RentalUsers").Append(&rentalUser)
	}
	return item, nil
}

// CreateLike likeを押す
func CreateLike(itemID, userID uint) (Item, error) {
	existed := false
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
	}
	db.Model(&item).Association("Likes").Append(&user)
	return item, nil
}

// CancelLike likeを消す
func CancelLike(itemID, userID uint) (Item, error) {
	existed := false
	item := Item{}
	db.Set("gorm:auto_preload", true).First(&item, itemID).Related(&item.Likes, "Likes")
	user, _ := GetUserByID(int(userID))
	for _, likeUser := range item.Likes {
		if likeUser.ID == userID {
			existed = true
		}
	}
	if !existed {
		return Item{}, errors.New("いいねしていません")
	}
	db.Model(&item).Association("Likes").Delete(&user)
	return item, nil
}

// SearchItems itemsをNameの部分一致で取得する
func SearchItems(searchString string) ([]Item, error) {
	res := []Item{}
	db.Where("name LIKE ?", "%"+searchString+"%").Find(&res)
	for i, item := range res {
		db.Set("gorm:auto_preload", true).First(&item).Related(&item.Owners, "Owners").Related(&item.Logs, "Logs").Related(&item.Likes, "Likes")
		var err error
		item.LatestLogs, err = GetLatestLogs(item.Logs)
		if err != nil {
			return []Item{}, err
		}
		item.LikeCounts = len(item.Likes)
		item.Likes = []User{}
		res[i] = item
	}
	return res, nil
}
