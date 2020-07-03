package model

// Tagmap tagmapの構造体
type Tagmap struct {
	GormModel
	ItemID int `gorm:"type:int;not null" json:"itemId"`
	TagID  int `gorm:"type:int;not null" json:"tagId"`
}

// TableName dbのテーブル名を指定する
func (tagmap *Tagmap) TableName() string {
	return "tagmaps"
}
