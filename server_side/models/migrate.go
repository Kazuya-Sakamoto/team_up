package models

import (
	//"time"

	"gopkg.in/gormigrate.v1"
	//"github.com/jinzhu/gorm"
)

var m *gormigrate.Gormigrate

func setMigrate() {
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{})
}
