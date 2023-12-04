package model

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

var allTables = []interface{}{
	User{},
	Item{},
	Log{},
	Owner{},
	Comment{},
	RentalUser{},
	File{},
}

type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// EstablishConnection DBに接続する
func EstablishConnection() (*gorm.DB, error) {
	user := os.Getenv("MYSQL_USERNAME")
	if user == "" {
		user = "root"
	}

	pass := os.Getenv("MYSQL_PASSWORD")
	if pass == "" {
		pass = ""
	}

	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}

	dbname := os.Getenv("MYSQL_DATABASE")
	if dbname == "" {
		dbname = "booq"
	}

	_db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)+"?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4")
	db = _db
	db.BlockGlobalUpdate(true)
	return db, err
}

// Migrate DBのマイグレーション
func Migrate() error {
	if err := db.AutoMigrate(allTables...).Error; err != nil {
		return err
	}

	traP, _ := GetUserByName("traP")
	if traP.Name == "" {
		user := User{
			Name:        "traP",
			DisplayName: "traP",
			Admin:       true,
		}
		_, err := CreateUser(user)
		if err != nil {
			return err
		}
	}

	sienka, _ := GetUserByName("sienka")
	if sienka.Name == "" {
		user := User{
			Name:        "sienka",
			DisplayName: "支援課",
			Admin:       true,
		}
		_, err := CreateUser(user)
		if err != nil {
			return err
		}
	}

	// https://cover.openbd.jp/xxx.jpg -> https://iss.ndl.go.jp/thumbnail/xxx
	db.Exec("UPDATE items SET img_url=REPLACE(REPLACE(img_url, '.jpg', ''), 'https://cover.openbd.jp/', 'https://iss.ndl.go.jp/thumbnail/') WHERE img_url LIKE 'https://cover.openbd.jp/%'")

	return nil
}
