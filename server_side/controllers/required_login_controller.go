package controllers

// RequiredLoginController ...
type RequiredLoginController struct {
	BaseController
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
