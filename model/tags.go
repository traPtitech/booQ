package model

// Tag tagの構造体
type Tag struct {
	GormModel
	Name string `gorm:"type:varchar(32);not null" json:"name"`
}

// TableName dbのテーブル名を指定する
func (tag *Tag) TableName() string {
	return "tags"
}
