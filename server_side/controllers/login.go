package controllers

import (
	//"fmt"
	"encoding/json"

	"app/server_side/models"

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
// @Param	body		body 	models.Staff	true		"body for staff content"
// @Success 200 {int} models.Staff.ID
// @Failure 403 body is empty
// @router / [post]
func (lc *LoginController) Post() {
	//パニックハンドリング
	if r := recover(); r != nil {
		lc.Ctx.ResponseWriter.WriteHeader(403)
	}

	session := lc.StartSession()
	staffID := session.Get("staffID")
	if staffID != nil {
		// UserID is not set, display another page
		lc.Ctx.WriteString("Already login.")
		return
	}

	// id/passwordを受け取る
	var requestStaff interface{}
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &requestStaff)
	if err != nil {
		lc.Ctx.ResponseWriter.WriteHeader(403)
		lc.Data["json"] = err.Error()
		lc.ServeJSON()
		return
	}

	requestLoginName := requestStaff.(map[string]interface{})["loginName"].(string)
	requestLoginPassword := requestStaff.(map[string]interface{})["loginPassword"].(string)
	// stuffデータを取得する
	targetStaff, err := models.GetUserByLoginName(requestLoginName)
	if err != nil {
		// staffの名前がまちがっていることをフロントに渡す
		//lc.Ctx.WriteString("ng")
		lc.Ctx.ResponseWriter.WriteHeader(401)
	}

	// パスワードの比較
	err = bcrypt.CompareHashAndPassword([]byte(targetStaff.LoginPassword), []byte(requestLoginPassword))
	if err != nil {
		//lc.Ctx.WriteString("ng")
		lc.Ctx.ResponseWriter.WriteHeader(401)
	} else {
		//lc.Ctx.WriteString("ok")
		session := lc.StartSession()
		session.Set("staffID", targetStaff.ID)
		lc.Data["json"] = targetStaff
		lc.ServeJSON()
	}
}

//GetOut ...
// @Title Logout
// @Description logout
// @Param	body		body 	models.Staff	true		"body for staff content"
// @Success 200 {int} models.Staff.ID
// @Failure 403 body is empty
// @router /logout [get]
func (lc *LoginController) GetOut() {
	//パニックハンドリング
	if r := recover(); r != nil {
		lc.Ctx.ResponseWriter.WriteHeader(403)
	}

	session := lc.StartSession()
	staffID := session.Get("staffID")
	if staffID == nil {
		// UserID is not set, display another page
		// TODO: change to login page(controller)
		lc.Ctx.WriteString("Already logout.")
		return
	}
	session.Delete("staffID")
	lc.Ctx.WriteString("Finished Logout.")
	return
}
