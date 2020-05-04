package models

import (
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"github.com/natefinch/lumberjack"

	// BrankImportForTes
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// DB servicesに渡すために定義
var DB *gorm.DB

// トランザクション周りを考慮して構成の変更を検討
func init() {
	initDB("teamup", false)
}

//InitTestDataBase ...
//For testmodel to use test database
func InitTestDataBase() *gorm.DB {
	initDB("test_db", true)
	return db
}

func initDB(dBName string, test bool) {
	// set default database
	host := beego.AppConfig.String("DBHost")
	// if host != "" {
	// 	host = "(" + host + ")"
	// }
	user := beego.AppConfig.String("DBUser")
	pass := beego.AppConfig.String("DBPassword")
	var err error
	// 全体のタイムゾーンを変更
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	sqlconn := beego.AppConfig.String("sqlconn")

	log.Println(sqlconn, user+":"+pass+"@"+host+"/"+dBName+"?charset=utf8mb4&parseTime=True&loc=Local")
	// 上記の変更によりloc=Localで良くなった
	db, err = gorm.Open(sqlconn, user+":"+pass+"@"+host+"/"+dBName+"?charset=utf8mb4&parseTime=True&loc=Local")
	DB = db
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	l := log.New(&lumberjack.Logger{
		Filename:   "./logs/testapp/db.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}, "", log.Ldate|log.Ltime)
	db.SetLogger(l)

	setMigrate()

	// db.DropTable("migrations")

	//if test {
	//  forTestInitMigrate()
	//} else {
	//暫定でテストでもこちらを行う

	initMigrate()
	//}
	if err = m.Migrate(); err != nil {
		log.Fatal(err)
	}
	log.Println("migration finished.")
}
