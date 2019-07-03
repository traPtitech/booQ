package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func EstablishConnection() (*gorm.DB, error) {
	user := os.Getenv("MYSQL_USERNAME")
	if user == "" {
		user = "root"
	}

	pass := os.Getenv("MYSQL_PASSWORD")
	if pass == "" {
		pass = "password"
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
