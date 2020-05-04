package controllers

import (
	"app/server_side/models"
	"encoding/json"
)

// RoleController ...
type RoleController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *RoleController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
}

//Post User
// @Title Post
// @Description create User
// @Param  body        body    models.User   true        "body for post content"
// @Success 201 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *RoleController) Post() {
	defer c.HandlePanic()
	var role models.Role
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &role)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	roleID, err := models.CreateRole(role)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		c.Data["json"] = map[string]int64{"roleId": roleID}
		c.Ctx.Output.SetStatus(201)
	}
	c.ServeJSON()
}

//GetAll Users
// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (c *RoleController) GetAll() {
	defer c.HandlePanic()
	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")
	roles, err := models.GetAllRoles(limit, offset)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		c.Data["json"] = roles
	}
	c.ServeJSON()
}

//Get User
// @Title Get
// @Description get User by UserID
// @Param  UserID        path    string  true        "The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :UserID is empty
// @router /:UserID [get]
func (c *RoleController) Get() {
	defer c.HandlePanic()
	roleID, err := c.GetInt64(":RoleID")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		role, err := models.GetRole(int64(roleID))
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(403)
		} else {
			c.Data["json"] = role
		}
	}
	c.ServeJSON()
}

//Put User
// @Title Update
// @Description update the User
// @Param  UserID        path    string  true        "The UserID you want to update"
// @Param  body        body    models.User   true        "body for User content"
// @Success 200 {object} models.User
// @Failure 403 :UserID is not int
// @router /:UserID [put]

//Delete User
// @Title Delete
// @Description delete the User
// @Param  UserID        path    string  true        "The UserID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 UserID is empty
// @router /:UserID [delete]
