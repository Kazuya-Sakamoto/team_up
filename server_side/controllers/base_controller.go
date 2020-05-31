package controllers

import (
	"app/server_side/errorcode"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// BaseController 共通ロジックの記述
type BaseController struct {
	beego.Controller
}

// エラーコードからエラーメッセージへの変換
func (bc *BaseController) postHandle(data interface{}, err error) {
	if err != nil {
		ei := getErrorInfoFromError(err)
		bc.Ctx.ResponseWriter.WriteHeader(ei.Code)
		bc.Data["json"] = ei.Message
	} else {
		bc.Data["json"] = data
		bc.Ctx.Output.SetStatus(201)
	}
}

func getErrorInfoFromError(err error) (ei errorcode.ErrorInfo) {
	logs.Error(err)
	es := strings.Split(err.Error(), ":::")
	if len(es) == 3 {
		ei = errorcode.GetErrorInfoFromErrorCode(es[1])
	} else {
		ei = errorcode.GetUnknownErrorInfo()
	}
	return
}

// エラーコードからエラーメッセージへの変換
func (bc *BaseController) getHandle(data interface{}, err error) {
	if err != nil {
		ei := getErrorInfoFromError(err)
		bc.Ctx.ResponseWriter.WriteHeader(ei.Code)
		bc.Data["json"] = ei.Message
	} else {
		bc.Data["json"] = data
	}
}

// エラーコードからエラーメッセージへの変換
func (bc *BaseController) putHandle(data interface{}, err error) {
	if err != nil {
		ei := getErrorInfoFromError(err)
		bc.Ctx.ResponseWriter.WriteHeader(ei.Code)
		bc.Data["json"] = ei.Message
	} else {
		if data == nil {
			bc.Data["json"] = "update success!"
		} else {
			bc.Data["json"] = data
		}
		bc.Ctx.Output.SetStatus(200)
	}
}

// エラーコードからエラーメッセージへの変換
func (bc *BaseController) deleteHandle(data interface{}, err error) {
	if err != nil {
		ei := getErrorInfoFromError(err)
		bc.Ctx.ResponseWriter.WriteHeader(ei.Code)
		bc.Data["json"] = ei.Message
	} else {
		if data == nil {
			bc.Data["json"] = "delete success!"
		} else {
			bc.Data["json"] = data
		}
		bc.Ctx.Output.SetStatus(200)
	}
}

func (bc *BaseController) parseErrorHandle(err error) {
	logs.Error(err)
	ei := errorcode.GetUnknownErrorInfo()
	bc.Ctx.ResponseWriter.WriteHeader(ei.Code)
	bc.Data["json"] = ei.Message
}

func (bc *BaseController) unmarshalErrorHandle(err error) {
	logs.Error(err)
	ei := errorcode.GetUnknownErrorInfo()
	bc.Ctx.ResponseWriter.WriteHeader(ei.Code)
	bc.Data["json"] = ei.Message
}

//HandlePanic ...
func (bc *BaseController) HandlePanic() {
	if err := recover(); err != nil {
		logs.Error(err)
		ei := errorcode.GetUnknownErrorInfo()
		bc.Data["json"] = ei.Message
		bc.Controller.Ctx.ResponseWriter.WriteHeader(ei.Code)
	}
}
