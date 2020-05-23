package controllers

// import (
// 	"app/server_side/models"
// 	"encoding/json"
// )

// // RoleController ...
// type RoleController struct {
// 	RequiredLoginController
// }

// // URLMapping ...
// func (c *RoleController) URLMapping() {
// 	c.Mapping("Get", c.Get)
// 	c.Mapping("GetAll", c.GetAll)
// }

// // Post ...
// // @Title Post
// // @Description create Role
// // @Param	body		body 	models.Role	true		"body for Role content"
// // @Success 201 {int} models.Role
// // @Failure 500 body is empty
// // @router / [post]
// func (c *RoleController) Post() {
// 	defer c.HandlePanic()
// 	var role models.Role
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &role)
// 	if err != nil {
// 		c.unmarshalErrorHandle(err)
// 		c.ServeJSON()
// 		return
// 	}
// 	c.postHandle(models.CreateRole(role))
// 	c.ServeJSON()
// }

// // Get Role
// // @Title Get
// // @Description get Role by RoleID
// // @Param	RoleID		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.Role
// // @Failure 500 :RoleID is empty
// // @router /:RoleID [get]
// func (c *RoleController) Get() {
// 	defer c.HandlePanic()
// 	roleID, err := c.GetInt64(":RoleID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		c.getHandle(models.GetRole(int64(roleID)))
// 	}
// 	c.ServeJSON()
// }

// // GetAll ...
// // @Title Get All
// // @Description get Role
// // @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// // @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// // @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// // @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// // @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// // @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// // @Success 200 {object} models.Role
// // @Failure 500
// // @router / [get]
// func (c *RoleController) GetAll() {
// 	defer c.HandlePanic()

// 	var limit, offset int64
// 	c.Ctx.Input.Bind(&limit, "limit")
// 	c.Ctx.Input.Bind(&offset, "offset")

// 	c.getHandle(models.GetAllRoles(limit, offset))
// 	c.ServeJSON()
// }

// // Put ...
// // @Title Put
// // @Description update the Role
// // @Param	RoleID		path 	string	true		"The RoleID you want to update"
// // @Param	body		body 	models.Role	true		"body for Role content"
// // @Success 200 {object} models.Role
// // @Failure 500 :RoleID is not int
// // @router /:RoleID [put]
// func (c *RoleController) Put() {
// 	defer c.HandlePanic()
// 	roleID, err := c.GetInt64(":RoleID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		var role models.Role
// 		err = json.Unmarshal(c.Ctx.Input.RequestBody, &role)
// 		if err != nil {
// 			c.unmarshalErrorHandle(err)
// 		} else {
// 			c.putHandle(nil, models.UpdateRole(int64(roleID), &role))
// 		}
// 	}
// 	c.ServeJSON()
// }

// //Delete Role
// // @Title Delete
// // @Description delete the Role
// // @Param  RoleID        path    string  true        "The RoleID you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 RoleID is empty
// // @router /:RoleID [delete]
// func (c *RoleController) Delete() {
// 	defer c.HandlePanic()
// 	roleID, err := c.GetInt64(":RoleID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		c.deleteHandle(nil, models.DeleteRole(int64(roleID)))
// 	}
// 	c.ServeJSON()
// }
