package controllers

import (
	"github.com/astaxie/beego"
)

// RequiredLoginController ...
type RequiredLoginController struct {
	beego.Controller
	UserID int64
}

// Prepare session
func (r *RequiredLoginController) Prepare() {
	defer r.HandlePanic()
	session := r.StartSession()
	userID := session.Get("userID")

	if userID == nil {
		r.Ctx.ResponseWriter.WriteHeader(401) // 開発中はコメントアウト
	} else {
		r.UserID = userID.(int64)
	}
}

// HandlePanic ...
func (r RequiredLoginController) HandlePanic() {
	if err := recover(); err != nil {
		r.Data["json"] = err
		r.Controller.Ctx.ResponseWriter.WriteHeader(403)
		r.ServeJSON()
	}
}
