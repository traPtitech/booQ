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
	db, err := EstablishConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = Migrate()
	if err != nil {
		panic(err)
	}
}
