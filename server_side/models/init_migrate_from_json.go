package models

import (
	//"database/sql"
	//"time"

	"encoding/json"
	"io/ioutil"

	//"github.com/amifiable-jp/jph-proto/dialects"
	// "github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

//Init ...
type Init struct {
	User []*User `json:"user"`
	Job  []*Job  `json:"job"`
	// AccessRight []*AccessRight `json:"accessRight"`
	// Role        []*Role        `json:"role"`
}

//BaseDirectory is project directory
// var BaseDirectory = os.Getenv("GOPATH") + "/src/app/server_side"
var BaseDirectory = "/go/src/app/server_side" //上記のGOPATHだとjsonのディレクトリがひろえなかったため変更

var initJSONDir = BaseDirectory + "/models/init_migrate_json"

func initMigrateFromJSON(tx *gorm.DB, subPath, jsonFileName string) (err error) {

	// if isDevServer := beego.AppConfig.String("isDevServer"); isDevServer == "true" {
	//  initJSONDir = "/var/www/destination/models/init_migrate_json"
	// }

	// if isProdServer := beego.AppConfig.String("isProdServer"); isProdServer == "true" {
	//  initJSONDir = "/var/www/current/dist/init_migrate_json"
	// }

	var init Init
	var initJSONByte []byte

	//JSONファイルを開く。
	initJSONByte, err = ioutil.ReadFile(initJSONDir + subPath + "/" + jsonFileName)
	if err != nil {
		return err
	}

	//JSONをパースする。
	json.Unmarshal(initJSONByte, &init)

	for _, m := range init.User {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.Job {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	// for _, m := range init.AccessRight {
	// 	if err = tx.Create(&m).Error; err != nil {
	// 		return err
	// 	}
	// }
	// for _, m := range init.Role {
	// 	if err = tx.Create(&m).Error; err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

