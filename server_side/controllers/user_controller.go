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

// // Post ...
// // @Title Post
// // @Description create User
// // @Param	body		body 	models.User	true		"body for User content"
// // @Success 201 {int} models.User
// // @Failure 500 body is empty
// // @router / [post]
// func (c *UserController) Post() {
// 	defer c.HandlePanic()
// 	var user models.User
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
// 	if err != nil {
// 		c.unmarshalErrorHandle(err)
// 		c.ServeJSON()
// 		return
// 	}
// 	c.postHandle(models.CreateUser(user))
// 	c.ServeJSON()
// }

// Get User
// @Title Get
// @Description get User by UserID
// @Param	UserID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 500 :UserID is empty
// @router /:UserID [get]
func (c *UserController) Get() {
	defer c.HandlePanic()
	userID, err := c.GetInt64(":UserID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetUser(int64(userID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 500
// @router / [get]
func (c *UserController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	c.getHandle(models.GetAllUsers(limit, offset))
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	UserID		path 	string	true		"The UserID you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 500 :UserID is not int
// @router /:UserID [put]
func (c *UserController) Put() {
	defer c.HandlePanic()
	userID, err := c.GetInt64(":UserID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		var user models.User
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &user)
		if err != nil {
			c.unmarshalErrorHandle(err)
		} else {
			c.putHandle(nil, models.UpdateUser(int64(userID), &user))
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
		c.parseErrorHandle(err)
	} else {
		c.deleteHandle(nil, models.DeleteUser(int64(userID)))
	}
	c.ServeJSON()
}
