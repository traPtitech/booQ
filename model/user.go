package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name       string    `gorm:"unique;not null;size:50" json:"name"`
	ScreenName string    `json:"screen_name"`
	IconFileID string    `json:"icon_file_id"`
	Group      string    `json:"group"`
	Admin      *int      `gorm:"default:0" json:"admin"`
}