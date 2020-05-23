package controllers

// import (
// 	"app/server_side/models"
// 	"encoding/json"
// )

// // AccessRightController ...
// type AccessRightController struct {
// 	RequiredLoginController
// }

// // URLMapping ...
// func (c *AccessRightController) URLMapping() {
// 	c.Mapping("Post", c.Post)
// 	c.Mapping("Get", c.Get)
// 	c.Mapping("GetAll", c.GetAll)
// 	c.Mapping("Put", c.Put)
// 	c.Mapping("Delete", c.Delete)
// }

// // Post ...
// // @Title Post
// // @Description create AccessRight
// // @Param	body		body 	models.AccessRight	true		"body for AccessRight content"
// // @Success 201 {int} models.AccessRight
// // @Failure 500 body is empty
// // @router / [post]
// func (c *AccessRightController) Post() {
// 	defer c.HandlePanic()
// 	var accessRight models.AccessRight
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &accessRight)
// 	if err != nil {
// 		c.unmarshalErrorHandle(err)
// 		c.ServeJSON()
// 		return
// 	}
// 	c.postHandle(models.CreateAccessRight(accessRight))
// 	c.ServeJSON()
// }

// // Get AccessRight
// // @Title Get
// // @Description get AccessRight by AccessRightID
// // @Param	AccessRightID		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.AccessRight
// // @Failure 500 :AccessRightID is empty
// // @router /:AccessRightID [get]
// func (c *AccessRightController) Get() {
// 	defer c.HandlePanic()
// 	accessRightID, err := c.GetInt64(":AccessRightID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		c.getHandle(models.GetAccessRight(int64(accessRightID)))
// 	}
// 	c.ServeJSON()
// }

// // GetAll ...
// // @Title Get All
// // @Description get AccessRight
// // @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// // @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// // @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// // @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// // @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// // @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// // @Success 200 {object} models.AccessRight
// // @Failure 500
// // @router / [get]
// func (c *AccessRightController) GetAll() {
// 	defer c.HandlePanic()

// 	var limit, offset int64
// 	c.Ctx.Input.Bind(&limit, "limit")
// 	c.Ctx.Input.Bind(&offset, "offset")

// 	c.getHandle(models.GetAllAccessRights(limit, offset))
// 	c.ServeJSON()
// }

// // Put ...
// // @Title Put
// // @Description update the AccessRight
// // @Param	AccessRightID		path 	string	true		"The AccessRightID you want to update"
// // @Param	body		body 	models.AccessRight	true		"body for AccessRight content"
// // @Success 200 {object} models.AccessRight
// // @Failure 500 :AccessRightID is not int
// // @router /:AccessRightID [put]
// func (c *AccessRightController) Put() {
// 	defer c.HandlePanic()
// 	accessRightID, err := c.GetInt64(":AccessRightID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		var accessRight models.AccessRight
// 		err = json.Unmarshal(c.Ctx.Input.RequestBody, &accessRight)
// 		if err != nil {
// 			c.unmarshalErrorHandle(err)
// 		} else {
// 			c.putHandle(nil, models.UpdateAccessRight(int64(accessRightID), &accessRight))
// 		}
// 	}
// 	c.ServeJSON()
// }

// //Delete AccessRight
// // @Title Delete
// // @Description delete the AccessRight
// // @Param  AccessRightID        path    string  true        "The AccessRightID you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 AccessRightID is empty
// // @router /:AccessRightID [delete]
// func (c *AccessRightController) Delete() {
// 	defer c.HandlePanic()
// 	accessRightID, err := c.GetInt64(":AccessRightID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		c.deleteHandle(nil, models.DeleteAccessRight(int64(accessRightID)))
// 	}
// 	c.ServeJSON()
// }
