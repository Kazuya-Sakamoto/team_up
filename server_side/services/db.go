package services

import (
	"app/server_side/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// func init() {
// 	initDB("teamup")
// }

// InitTestDataBase ...
func InitTestDataBase() *gorm.DB {
	db = models.InitTestDataBase()
	return db
}

// func initDB(dbName string) {
// 	db = models.DB
// }
