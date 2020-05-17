package models

import "time"

// User ...
type User struct {
	Model
	LoginName         string    `gorm:"" json:"loginName"`         // ログイン名
	LoginPassword     string    `gorm:"" json:"loginPassword"`     // ログインパスワード
	UserName          string    `gorm:"" json:"userName"`          // ユーザー名
	UserBirthday      time.Time `gorm:"" json:"userBirthday"`      // 生年月日
	Bio               string    `gorm:"" json:"bio"`               // 自己紹介文
	GithubAccount     string    `gorm:"" json:"githubAccount"`     // Githubアカウント
	TwitterAccount    string    `gorm:"" json:"twitterAccount"`    // Twitterアカウント
	LearningStartDate time.Time `gorm:"" json:"learningStartDate"` // 学習開始日

	// RoleID        int64  `gorm:"" json:"roleId"`
	// Role          *Role  `gorm:"" json:"role"`
}

// CreateUser ...
func CreateUser(user User) (UserID int64, err error) {
	err = db.Create(&user).Error
	return user.ID, err
}

// GetUser ...
func GetUser(UserID int64) (user User, err error) {
	err = db.Set("gorm:auto_preload", true).First(&user, UserID).Error
	return user, err
}

// GetUserByLoginName ...
func GetUserByLoginName(LoginName string) (user *User, err error) {
	user = &User{}
	err = db.Set("gorm:auto_preload", true).First(&user, User{LoginName: LoginName}).Error
	return user, err
}

// GetAllUsers ...
func GetAllUsers(limit int64, offset int64) (ml []*User, err error) {
	tx := db.Set("gorm:auto_preload", true).Begin()

	if limit != 0 {
		tx = tx.Limit(limit)
	} else {
		var count int64
		tx.Model(&ml).Count(&count)
		tx = tx.Limit(count)
	}

	// 関連テーブルを抽出条件に含める例
	// tx = tx.
	// 	Joins("JOIN roles ON roles.id = role_id").
	// 	Joins("JOIN role_access_rights ON role_access_rights.role_id = roles.id").
	// 	Joins("JOIN access_rights ON role_access_rights.access_right_id = access_rights.id").
	// 	Where("access_rights.access_right_name = ?", "MOP")

	err = tx.Offset(offset).Find(&ml).Commit().Error

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
