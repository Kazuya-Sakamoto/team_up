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
	User                []*User                `json:"user"`
	Job                 []*Job                 `json:"job"`
	UserAuthInfo        []*UserAuthInfo        `json:"userAuthInfo"`
	ProgramingLanguage  []*ProgramingLanguage  `json:"programingLanguage"`
	ProgramingFramework []*ProgramingFramework `json:"programingFramework"`
	Skill               []*Skill               `json:"skill"`
	CommunicationTool   []*CommunicationTool   `json:"communicationTool"`
	PositionTag         []*PositionTag         `json:"positionTag"`
	JobStatus           []*JobStatus           `json:"jobStatus"`
	IndividualPortfolio []*IndividualPortfolio `json:"individualPortfolio"`
	ChatMessage         []*ChatMessage         `json:"chatMessage"`
	FavoriteJob         []*FavoriteJob         `json:"favoriteJob"`
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
	for _, m := range init.UserAuthInfo {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.ProgramingLanguage {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.ProgramingFramework {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.Skill {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.CommunicationTool {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.PositionTag {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.JobStatus {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.IndividualPortfolio {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.ChatMessage {
		if err = tx.Create(&m).Error; err != nil {
			return err
		}
	}
	for _, m := range init.FavoriteJob {
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
