package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

var allTables = []interface{}{
	User{},
	Item{},
	Owner{},
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

	dbname := os.Getenv("MYSQL_DATABASE")
	if dbname == "" {
		dbname = "booq"
	}

	_db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, dbname)+"?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4")
	db = _db
	return db, err
}

// Migrate DBのマイグレーション
func Migrate() error {
	if err := db.AutoMigrate(allTables...).Error; err != nil {
		return err
	}

	traP, _ := GetUser(User{Name: "traP"})
	if traP.Name == "" {
		user := User{
			Name:        "traP",
			DisplayName: "traP",
			IconFileID:  "099eed74-3ab3-4655-ac37-bc7df1139b3d",
			Admin:       true,
		}
		_, _ = CreateUser(user)
	}

	sienka, err := GetUser(User{Name: "sienka"})
	if sienka.Name == "" {
		user := User{
			Name:        "sienka",
			DisplayName: "支援課",
			IconFileID:  "",
			Admin:       true,
		}
		_, err = CreateUser(user)
	}

	return err
}
