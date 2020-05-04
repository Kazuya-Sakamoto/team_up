package controllers

import (
	"encoding/json"

	"app/server_side/models"
	// "github.com/astaxie/beego"
)

//UserController Operations
type UserController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

//Post User
// @Title Post
// @Description create User
// @Param  body        body    models.User   true        "body for post content"
// @Success 201 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	defer c.HandlePanic()
	var user models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	userID, err := models.CreateUser(user)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		c.Data["json"] = map[string]int64{"userId": userID}
		c.Ctx.Output.SetStatus(201)
	}
	c.ServeJSON()
}

//GetAll Users
// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (c *UserController) GetAll() {
	defer c.HandlePanic()
	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")
	users, err := models.GetAllUsers(limit, offset)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		c.Data["json"] = users
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
func (c *UserController) Get() {
	defer c.HandlePanic()
	userID, err := c.GetInt64(":UserID")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		user, err := models.GetUser(int64(userID))
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(403)
		} else {
			c.Data["json"] = user
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
func (c *UserController) Put() {
	defer c.HandlePanic()
	userID, err := c.GetInt64(":UserID")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		var user models.User
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &user)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}
		err = models.UpdateUser(int64(userID), &user)
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(403)
		} else {
			c.Data["json"] = "update success!"
		}
	}
	c.ServeJSON()
}

//Delete User
// @Title Delete
// @Description delete the User
// @Param  UserID        path    string  true        "The UserID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 UserID is empty
// @router /:UserID [delete]
func (c *UserController) Delete() {
	defer c.HandlePanic()
	userID, err := c.GetInt64(":UserID")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		err = models.DeleteUser(int64(userID))
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(403)
		} else {
			c.Data["json"] = "delete success!"
		}
	}
	c.ServeJSON()
}
