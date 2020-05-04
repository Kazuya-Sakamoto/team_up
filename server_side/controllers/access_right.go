package controllers

import (
	"app/server_side/models"
	"encoding/json"
)

// AccessRightController ...
type AccessRightController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *AccessRightController) URLMapping() {
	// c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	// c.Mapping("Delete", c.Delete)
}

//Post AccessRight
// @Title Post
// @Description create AccessRight
// @Param  body        body    models.AccessRight   true        "body for post content"
// @Success 201 {int} models.AccessRight.Id
// @Failure 403 body is empty
// @router / [post]
func (c *AccessRightController) Post() {
	defer c.HandlePanic()
	var accessRight models.AccessRight
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &accessRight)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	accessRightID, err := models.CreateAccessRight(accessRight)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		c.Data["json"] = map[string]int64{"accessRightId": accessRightID}
		c.Ctx.Output.SetStatus(201)
	}
	c.ServeJSON()
}

//GetAll AccessRights
// @Title GetAll
// @Description get all AccessRights
// @Success 200 {object} models.AccessRight
// @router / [get]
func (c *AccessRightController) GetAll() {
	defer c.HandlePanic()
	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")
	accessRights, err := models.GetAllAccessRights(limit, offset)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		c.Data["json"] = accessRights
	}
	c.ServeJSON()
}

//Get AccessRight
// @Title Get
// @Description get AccessRight by AccessRightID
// @Param  AccessRightID        path    string  true        "The key for staticblock"
// @Success 200 {object} models.AccessRight
// @Failure 403 :AccessRightID is empty
// @router /:AccessRightID [get]
func (c *AccessRightController) Get() {
	defer c.HandlePanic()
	accessRightID, err := c.GetInt64(":AccessRightID")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		accessRight, err := models.GetAccessRight(int64(accessRightID))
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(403)
		} else {
			c.Data["json"] = accessRight
		}
	}
	c.ServeJSON()
}

//Put AccessRight
// @Title Update
// @Description update the AccessRight
// @Param  AccessRightID        path    string  true        "The AccessRightID you want to update"
// @Param  body        body    models.AccessRight   true        "body for AccessRight content"
// @Success 200 {object} models.AccessRight
// @Failure 403 :AccessRightID is not int
// @router /:AccessRightID [put]

//Delete AccessRight
// @Title Delete
// @Description delete the AccessRight
// @Param  AccessRightID        path    string  true        "The AccessRightID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 AccessRightID is empty
// @router /:AccessRightID [delete]
