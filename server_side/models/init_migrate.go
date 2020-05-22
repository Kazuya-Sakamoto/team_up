package models

import (
	//"database/sql"
	//"time"

	"io/ioutil"
	"strings"

	// "time"

	// "app/sampleapp/dialects"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

func initMigrateFromJSONDir(tx *gorm.DB, path string) error {
	//そのディレクトリを読み込む。
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	//それぞれの内容物に関して、
	for _, content := range dir {
		//それがディレクトリならば、
		if content.IsDir() == true {
			//ディレクトリ階層を記録する。
			path = path + "/" + content.Name()

			//再帰。横優先探索のため、defer。
			initMigrateFromJSONDir(tx, path)
		} else { //それがファイルならば、
			fileName := content.Name() //ファイル名

			//pathからinitJSONDirの文字列を削除する
			subPath := strings.Replace(path, initJSONDir, "", -1)

			//イニットマイグレーション
			err := initMigrateFromJSON(tx, subPath, fileName)
			if err != nil {
				//エラー処理
				return err
			}
		}
	}

	return nil
}

func initMigrate() {
	m.InitSchema(func(tx *gorm.DB) error {
		// add tables
		err := tx.AutoMigrate(
			&User{},
			&Job{},
			&UserAuthInfo{},
			// &AccessRight{},
			// &Role{},

			// //History Tables
			// &StaffHistory{},

		).Error
		if err != nil {
			return err
		}

		//init_migrate_jsonフォルダを読み込む。
		if isDevServer := beego.AppConfig.String("isDevServer"); isDevServer == "true" {
			initJSONDir = "/var/www/destination/models/init_migrate_json"
		}

		if isProdServer := beego.AppConfig.String("isProdServer"); isProdServer == "true" {
			initJSONDir = "/var/www/current/dist/init_migrate_json"
		}

		err = initMigrateFromJSONDir(tx, initJSONDir)
		if err != nil {
			beego.Debug(err)
			logs.Error(err)
		}

		// // Add ForeignKey
		// if err := tx.Model(&User{}).AddForeignKey("gender_type_id", "gender_types(id)", "RESTRICT", "CASCADE").Error; err != nil {
		// 	return err
		// }

		return nil
	})
}
