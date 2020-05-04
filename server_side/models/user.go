package models

// User ...
type User struct {
	Model
	UserName      string `gorm:"" json:"userName"`
	LoginName     string `gorm:"" json:"loginName"`
	LoginPassword string `gorm:"" json:"loginPassword"`
}

// CreateUser ...
func CreateUser(user User) (UserID int64, err error) {
	err = db.Create(&user).Error
	return user.ID, err
}

// GetUser ...
func GetUser(UserID int64) (user User, err error) {
	err = db.First(&user, UserID).Error
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
	tx := db.Begin()

	if limit != 0 {
		tx = tx.Limit(limit)
	} else {
		var count int64
		db.Model(&ml).Count(&count)
		tx = tx.Limit(count)
	}

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
