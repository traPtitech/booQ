package model

import (
	"os"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup() {
	_, err := EstablishConnection()
	if err != nil {
		panic(err)
	}

	err = Migrate()
	if err != nil {
		panic(err)
	}
	user := User{
		Name:        "testUser",
		DisplayName: "テストユーザー",
		Admin:       false,
	}
	_, err = CreateUser(user)
	if err != nil {
		panic(err)
	}
}
