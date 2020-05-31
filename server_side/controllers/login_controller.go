package controllers

import (
	"app/server_side/models"
	"app/server_side/services"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

//LoginController Operations
type LoginController struct {
	beego.Controller
}

//Post Login
// @Title Login
// @Description login
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.ID
// @Failure 403 body is empty
// @router / [post]
func (lc *LoginController) Post() {
	if r := recover(); r != nil {
		lc.Ctx.ResponseWriter.WriteHeader(403)
	}

	session := lc.StartSession()
	userID := session.Get("userID")
	log.Println("Session UserID", userID)
	if userID != nil {
		lc.Ctx.WriteString("Already login.")
		return
	}

	// id/passwordを受け取る
	var userAuthInfo models.UserAuthInfo
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &userAuthInfo)
	log.Println(userAuthInfo)
	if err != nil {
		lc.Ctx.ResponseWriter.WriteHeader(403)
		lc.Data["json"] = err.Error()
		lc.ServeJSON()
		return
	}

	requestLoginName := userAuthInfo.LoginName
	requestLoginPassword := userAuthInfo.LoginPassword
	log.Println("requestLoginName", requestLoginName)
	log.Println("requestLoginPassword", requestLoginPassword)

	targetUser, err := services.PostMethodLogin(requestLoginName)
	log.Println("targetUser", targetUser)
	if err != nil {
		// userの名前がまちがっていることをフロントに渡す
		lc.Ctx.ResponseWriter.WriteHeader(401)
	}

	// パスワードの比較
	err = bcrypt.CompareHashAndPassword([]byte(targetUser.LoginPassword), []byte(requestLoginPassword))
	if err != nil {
		lc.Ctx.ResponseWriter.WriteHeader(401)
	} else {
		// session := lc.StartSession()
		session.Set("userID", targetUser.ID)
		lc.Data["json"] = targetUser
		lc.ServeJSON()
	}
}

//GetOut ...
// @Title Logout
// @Description logout
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.ID
// @Failure 403 body is empty
// @router /logout [get]
func (lc *LoginController) GetOut() {
	//パニックハンドリング
	if r := recover(); r != nil {
		lc.Ctx.ResponseWriter.WriteHeader(403)
	}

	session := lc.StartSession()
	userID := session.Get("userID")
	if userID == nil {
		// UserID is not set, display another page
		// TODO: change to login page(controller)
		lc.Ctx.WriteString("Already logout.")
		return
	}
	session.Delete("userID")
	lc.Ctx.WriteString("Finished Logout.")
	return
}
