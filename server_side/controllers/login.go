package controllers

import (
	//"fmt"
	"app/server_side/services"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/session"
)

//LoginController Operations
type LoginController struct {
	beego.Controller
	// session.Store
}

//Post Login
// @Title Login
// @Description login
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.ID
// @Failure 403 body is empty
// @router / [post]
func (lc *LoginController) Post() {
	//パニックハンドリング
	if r := recover(); r != nil {
		lc.Ctx.ResponseWriter.WriteHeader(403)
	}

	session := lc.StartSession()
	userID := session.Get("userID")
	if userID != nil {
		// UserID is not set, display another page
		lc.Ctx.WriteString("Already login.")
		return
	}

	// id/passwordを受け取る
	var requestUser interface{}
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &requestUser)
	if err != nil {
		lc.Ctx.ResponseWriter.WriteHeader(403)
		lc.Data["json"] = err.Error()
		lc.ServeJSON()
		return
	}

	requestLoginName := requestUser.(map[string]interface{})["loginName"].(string)
	requestLoginPassword := requestUser.(map[string]interface{})["loginPassword"].(string)
	// userデータを取得する
	targetUser, err := services.PostMethodLogin(requestLoginName)
	if err != nil {
		// userの名前がまちがっていることをフロントに渡す
		//lc.Ctx.WriteString("ng")
		lc.Ctx.ResponseWriter.WriteHeader(401)
	}

	// パスワードの比較
	err = bcrypt.CompareHashAndPassword([]byte(targetUser.LoginPassword), []byte(requestLoginPassword))
	if err != nil {
		//lc.Ctx.WriteString("ng")
		lc.Ctx.ResponseWriter.WriteHeader(401)
	} else {
		//lc.Ctx.WriteString("ok")
		session := lc.StartSession()
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
