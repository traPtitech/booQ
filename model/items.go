package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Item itemの構造体
type Item struct {
	gorm.Model
	Name        string        `gorm:"type:varchar(64);not null" json:"name"`
	Type        int           `gorm:"type:int;not null" json:"type"`
	Code        string        `gorm:"type:varchar(13);" json:"code"`
	Description string        `gorm:"type:text;" json:"description"`
	ImgURL      string        `gorm:"type:text;" json:"img_url"`
	Owners      []*Owner      `gorm:"many2many:ownership_maps;" json:"owners"`
	RentalUsers []*RentalUser `gorm:"many2many:rental_user_maps;" json:"rental_users"`
	Logs        []Log         `json:"logs"`
	LatestLogs  []Log         `json:"latest_logs"`
	Comments    []Comment     `json:"comments"`
	Likes       []User        `gorm:"many2many:like_maps;" json:"likes"`
	LikeCounts  int           `gorm:"-" json:"like_counts"`
}

type Owner struct {
	gorm.Model
	UserID     uint `gorm:"type:int;not null" json:"owner_id"`
	User       User `json:"user"`
	Rentalable bool `gorm:"type:bool;not null;default:true" json:"rentalable"`
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
	db.Set("gorm:auto_preload", true).Preload("Owners.User").Preload("Logs.User").Preload("RentalUsers.User").Preload("Comments.User").First(&res, id)
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
	db.Set("gorm:auto_preload", true).Preload("Owners.User").Preload("Logs.User").Preload("RentalUsers.User").Preload("Comments.User").First(&res, "name = ?", name)
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
	db.Set("gorm:auto_preload", true).Preload("Owners.User").Preload("Logs.User").Preload("RentalUsers.User").Preload("Comments.User").Find(&res)
	for i, item := range res {
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
	db.Set("gorm:auto_preload", true).Find(&item).Related(&item.Owners, "Owners")
	owner.User, _ = GetUserByID(int(owner.UserID))
	for _, nowOwner := range item.Owners {
		if nowOwner.UserID != owner.UserID {
			continue
		}
		existed = true
		return Item{}, errors.New("該当の物品をすでに所有しています")
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

func AddOwner(owner Owner, item Item) (Item, error) {
	var existed bool
	db.Preload("Owners").Preload("Logs").Find(&item)
	latestLog, err := GetLatestLog(item.Logs, owner.UserID)
	if err != nil {
		return Item{}, err
	}
	log := Log{}
	rentalableCount := latestLog.Count
	owner.User, _ = GetUserByID(int(owner.UserID))
	for i, nowOwner := range item.Owners {
		if nowOwner.UserID != owner.UserID {
			continue
		}
		nowOwner.Rentalable = owner.Rentalable
		if owner.Count-nowOwner.Count+rentalableCount < 0 {
			return Item{}, errors.New("現在貸し出し中の物品が存在するのでそれよりも少ない数にはできません")
		}
		if owner.Count-nowOwner.Count < 0 {
			log.Type = 3
		} else {
			log.Type = 2
		}
		log.Count = owner.Count - nowOwner.Count
		nowOwner.Count = owner.Count
		existed = true
		nowOwner.User = owner.User
		db.Save(&nowOwner)
		item.Owners[i] = nowOwner
		latestLog, err := GetLatestLog(item.Logs, owner.UserID)
		if err != nil {
			return Item{}, err
		}
		log.ItemID = latestLog.ItemID
		log.UserID = owner.UserID
		log.OwnerID = owner.UserID
		log.Count += latestLog.Count
		if latestLog.ItemID != 0 {
			_, err = CreateLog(log)
			if err != nil {
				return Item{}, err
			}
		}
	}
	if !existed {
		return Item{}, errors.New("該当の物品を所有していないため変更できません")
	}
	return item, nil
}

// RentalItem 物品を借りたりするときにRentalUserを作成する
func RentalItem(rentalUser RentalUser, item Item) (Item, error) {
	var existed bool
	db.Set("gorm:auto_preload", true).Preload("Logs.User").Preload("RentalUsers.User").Preload("Comments.User").Find(&item)
	// owner.User, _ = GetUserByID(int(owner.UserID))
	for i, nowRentalUser := range item.RentalUsers {
		if nowRentalUser.UserID != rentalUser.UserID || nowRentalUser.OwnerID != rentalUser.OwnerID {
			continue
		}
		existed = true
		nowRentalUser.Count += rentalUser.Count
		if nowRentalUser.Count > 0 {
			return Item{}, errors.New("Return超過3")
		}
		db.Save(&nowRentalUser)
		item.RentalUsers[i] = nowRentalUser
	}
	if !existed {
		if rentalUser.Count > 0 {
			return Item{}, errors.New("該当のUserは指定のItemを借りていません")
		}
		db.Create(&rentalUser)
		db.Model(&item).Association("RentalUsers").Append(&rentalUser)
	}
	return item, nil
}

// CreateLike likeを押す
func CreateLike(itemID, userID uint) (Item, error) {
	existed := false
	item := Item{}
	db.Set("gorm:auto_preload", true).First(&item, itemID)
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
	db.Set("gorm:auto_preload", true).First(&item, itemID)
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

// SearchItemsByOwner itemsをOwnerNameから取得する
func SearchItemByOwner(ownerName string) ([]Item, error) {
	res := []Item{}
	items := []Item{}
	owner, err := GetUserByName(ownerName)
	if owner.ID == 0 {
		return []Item{}, errors.New("該当のUserが存在しません")
	}
	if err != nil {
		return []Item{}, err
	}
	db.Set("gorm:auto_preload", true).Preload("Logs.User").Preload("RentalUsers.User").Preload("Comments.User").Preload("Owners.User").Find(&res)
	for _, item := range res {
		var err error
		item.LatestLogs, err = GetLatestLogs(item.Logs)
		if err != nil {
			return []Item{}, err
		}
		for _, owner := range item.Owners {
			if owner.User.Name == ownerName {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

// SearchItemsByRental itemsをRentalUserNameから取得する
func SearchItemByRental(rentalUserID uint) ([]Item, error) {
	items := []Item{}
	res := []Item{}
	db.Set("gorm:auto_preload", true).Preload("Logs.User").Preload("RentalUsers.User").Preload("Comments.User").Find(&items)
	for _, item := range items {
		var err error
		item.LatestLogs, err = GetLatestLogs(item.Logs)
		if err != nil {
			return []Item{}, err
		}
		for _, rentalUser := range item.RentalUsers {
			if rentalUser.UserID == rentalUserID && rentalUser.Count < 0 {
				res = append(res, item)
			}
		}
	}
	return res, nil
}

// SearchItems itemsをNameの部分一致で取得する
func SearchItems(searchString string) ([]Item, error) {
	res := []Item{}
	db.Set("gorm:auto_preload", true).Preload("Logs.User").Preload("Owners.User").Preload("RentalUsers.User").Preload("Comments.User").Where("name LIKE ?", "%"+searchString+"%").Find(&res)
	for i, item := range res {
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

// DestroyItem itemを削除する
func DestroyItem(item Item) (Item, error) {
	// ここでは指定のItemがあるかどうか判定していません
	db.Delete(&item)
	return item, nil
}

// UpdateItem itemを変更する
func UpdateItem(item *Item, body *map[string]interface{}, isAdmin bool) Item {
	fields := []string{"name", "code", "description", "img_url"}
	if isAdmin {
		fields = append(fields, "type")
	}
	db.Model(item).Updates(filterMap(body, fields))

	return *item
}

func filterMap(input *map[string]interface{}, keys []string) map[string]interface{} {
	output := make(map[string]interface{})
	for _, key := range keys {
		if val, ok := (*input)[key]; ok {
			output[key] = val
		}
	}
	return output
}
