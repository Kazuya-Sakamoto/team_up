package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// UserAuthInfo スタッフ認証情報テーブル
type UserAuthInfo struct {
	Model
	UserID        int64  `gorm:"" json:"userId"`        // スタッフID
	User          *User  `gorm:"" json:"user"`          // スタッフ情報
	LoginName     string `gorm:"" json:"loginName"`     // ログイン名
	LoginPassword string `gorm:"" json:"loginPassword"` // ログインパスワード（ハッシュ化済み）
}

func init() {
}

// CreateUserAuthInfo ...
func CreateUserAuthInfo(tx *gorm.DB, userAuthInfo UserAuthInfo) (UserID int64, err error) {
	if userAuthInfo.LoginName == "" {
		err = fmt.Errorf("user's login name must not be empty")
		return 0, err
	}
	if userAuthInfo.LoginPassword == "" {
		err = fmt.Errorf("user's login password must not be empty")
		return 0, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(userAuthInfo.LoginPassword), 10)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	userAuthInfo.LoginPassword = string(hash)
	err = tx.Create(&userAuthInfo).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return 0, err
	}
	return
}

// // CreateUserAuthInfoFromUserApprovalAuthInfo 注:UserApprovalAuthInfoからのコピー用途専用
// func CreateUserAuthInfoFromUserApprovalAuthInfo(tx *gorm.DB, userAuthInfo UserAuthInfo) (UserID int64, err error) {
// 	err = tx.Create(&userAuthInfo).Error
// 	if err != nil {
// 		logs.Error(err)
// 		tx.Rollback()
// 		return 0, err
// 	}
// 	return
// }

// FirstUserAuthInfoWithID ...
func FirstUserAuthInfoWithID(tx *gorm.DB, userAuthInfoID int64) (userAuthInfo UserAuthInfo, err error) {
	err = tx.Set("gorm:auto_preload", true).First(&userAuthInfo, userAuthInfoID).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return
}

// FirstUserAuthInfoWithLoginName ...
func FirstUserAuthInfoWithLoginName(tx *gorm.DB, loginName string) (userAuthInfo UserAuthInfo, err error) {
	userAuthInfo = UserAuthInfo{}
	err = tx.Set("gorm:auto_preload", true).First(&userAuthInfo, UserAuthInfo{LoginName: loginName}).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return
}

// FindUserAuthInfoWithUserID ...
func FindUserAuthInfoWithUserID(tx *gorm.DB, userID int64) (userAuthInfo UserAuthInfo, err error) {
	err = tx.Set("gorm:auto_preload", true).First(&userAuthInfo, userID).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return
}

// UpdateUserAuthInfo ...
func UpdateUserAuthInfo(tx *gorm.DB, userAuthInfoID int64, userAuthInfo *UserAuthInfo) (err error) {
	if userAuthInfo.LoginPassword != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(userAuthInfo.LoginPassword), 10)
		if err != nil {
			logs.Error(err)
			return err
		}
		userAuthInfo.LoginPassword = string(hash)
	}
	err = tx.Model(&UserAuthInfo{Model: Model{ID: userAuthInfoID}}).Update(userAuthInfo).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return
}

// // UpdateUserAuthInfoFromUserApprovalAuthInfo 注:UserApprovalAuthInfoからのコピー用途専用
// func UpdateUserAuthInfoFromUserApprovalAuthInfo(tx *gorm.DB, userAuthInfoID int64, userAuthInfo *UserAuthInfo) (err error) {
// 	if userAuthInfo.LoginPassword != "" {
// 		hash, err := bcrypt.GenerateFromPassword([]byte(userAuthInfo.LoginPassword), 10)
// 		if err != nil {
// 			logs.Error(err)
// 			return err
// 		}
// 		userAuthInfo.LoginPassword = string(hash)
// 	}
// 	err = tx.Model(&UserAuthInfo{Model: Model{ID: userAuthInfoID}}).Update(userAuthInfo).Error
// 	if err != nil {
// 		logs.Error(err)
// 		tx.Rollback()
// 		return
// 	}
// 	return
// }

// DeleteUserAuthInfo ...
func DeleteUserAuthInfo(tx *gorm.DB, userAuthInfoID int64) (err error) {
	err = tx.Delete(&UserAuthInfo{Model: Model{ID: userAuthInfoID}}).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return
}
