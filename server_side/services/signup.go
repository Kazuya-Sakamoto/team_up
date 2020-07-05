package services

import (
	"app/server_side/models"

	"github.com/astaxie/beego/logs"
)

// PostMethodSignup SignupController POST method
func PostMethodSignup(userAuthInfo models.UserAuthInfo) (userID int64, err error) {
	tx := db.Begin()
	user := models.User{UserName: "No Name"}
	userID, err = models.CreateUser(tx, user)
	if err != nil {
		logs.Error(err)
		return
	}
	userAuthInfo.UserID = userID
	_, err = models.CreateUserAuthInfo(tx, userAuthInfo)
	if err != nil {
		logs.Error(err)
		return
	}
	err = tx.Commit().Error
	return
}
