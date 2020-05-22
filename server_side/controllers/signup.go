package controllers

import (
	"app/server_side/models"
	"app/server_side/services"
	"encoding/json"

	"github.com/astaxie/beego"
)

// SignupController ...
type SignupController struct {
	beego.Controller
}

//Post Signup
// @Title Signup
// @Description Signup
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.ID
// @Failure 403 body is empty
// @router / [post]
func (lc *SignupController) Post() {
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
	var userAuthInfo models.UserAuthInfo
	err := json.Unmarshal(lc.Ctx.Input.RequestBody, &userAuthInfo)
	if err != nil {
		lc.Ctx.ResponseWriter.WriteHeader(403)
		lc.Data["json"] = err.Error()
		lc.ServeJSON()
		return
	}

	userID, err = services.PostMethodSignup(userAuthInfo)
	if err != nil {
		lc.Data["json"] = err.Error()
		lc.Ctx.ResponseWriter.WriteHeader(401)
	} else {
		session := lc.StartSession()
		session.Set("userID", userID)
		lc.Data["json"] = map[string]int64{"userId": userID.(int64)} // userIDはinterface形のためuserID.(int64)とすることでint64にCast（型変換）
		lc.Ctx.Output.SetStatus(201)
	}
	lc.ServeJSON()
}
