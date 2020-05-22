package services

import (
	"app/server_side/models"

	"github.com/astaxie/beego/logs"
)

// PostMethodLogin Login Controller POST method
func PostMethodLogin(loginName string) (userAuthInfo models.UserAuthInfo, err error) {
	tx := db.Begin()
	userAuthInfo, err = models.FirstUserAuthInfoWithLoginName(tx, loginName)
	if err != nil {
		logs.Error(err)
		return
	}
	err = tx.Commit().Error
	return
}
