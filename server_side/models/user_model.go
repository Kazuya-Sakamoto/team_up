package models

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	Model
	UserName          string     `gorm:"" json:"userName"`          // ユーザー名
	UserBirthday      *time.Time `gorm:"" json:"userBirthday"`      // 生年月日
	Bio               string     `gorm:"" json:"bio"`               // 自己紹介文
	GithubAccount     string     `gorm:"" json:"githubAccount"`     // Githubアカウント
	TwitterAccount    string     `gorm:"" json:"twitterAccount"`    // Twitterアカウント
	LearningStartDate *time.Time `gorm:"" json:"learningStartDate"` // 学習開始日
	Job               []*Job     `gorm:"PRELOAD:false" json:"job"`  // 案件

	// RoleID        int64  `gorm:"" json:"roleId"`
	// Role          *Role  `gorm:"" json:"role"`
}

// CreateUser ...
func CreateUser(tx *gorm.DB, user User) (UserID int64, err error) {
	err = tx.Create(&user).Error
	if err != nil {
		logs.Error("models", err)
		tx.Rollback()
		return 0, err
	}
	return user.ID, err
}

// GetUser ...
func GetUser(UserID int64) (user User, err error) {
	err = db.Set("gorm:auto_preload", true).First(&user, UserID).Error
	return user, err
}

// GetAllUsers ...
func GetAllUsers(limit int64, offset int64) (ml []*User, err error) {
	tx := db.Preload("Job").Begin()

	if limit != 0 {
		tx = tx.Limit(limit)
	} else {
		var count int64
		tx.Model(&ml).Count(&count)
		tx = tx.Limit(count)
	}

	err = tx.Offset(offset).Find(&ml).Commit().Error

	// 関連テーブルの抽出条件設定例
	// tx = tx.
	// Joins("JOIN roles ON roles.id = role_id").
	// Joins("JOIN role_access_rights ON role_access_rights.role_id = roles.id").
	// Joins("JOIN access_rights ON role_access_rights.access_right_id = access_rights.id").
	// Where("access_rights.access_right_name = ?", "MOP")

	return ml, err
}

// UpdateUser ...
func UpdateUser(UserID int64, user *User) (err error) {
	err = db.Model(&User{Model: Model{ID: UserID}}).Update(user).Error
	return err
}

// DeleteUser ...
func DeleteUser(UserID int64) (err error) {
	err = db.Delete(&User{Model: Model{ID: UserID}}).Error
	return err
}
