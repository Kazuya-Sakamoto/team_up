package controllers

// import (
// 	"app/server_side/models"
// 	"encoding/json"
// )

// // PositionTagController ...
// type PositionTagController struct {
// 	RequiredLoginController
// }

// // URLMapping ...
// func (c *PositionTagController) URLMapping() {
// 	c.Mapping("Post", c.Post)
// 	c.Mapping("Get", c.Get)
// 	c.Mapping("GetAll", c.GetAll)
// 	c.Mapping("Put", c.Put)
// 	c.Mapping("Delete", c.Delete)
// }

// // Post ...
// // @Title Post
// // @Description create PositionTag
// // @Param	body		body 	models.PositionTag	true		"body for PositionTag content"
// // @Success 201 {int} models.PositionTag
// // @Failure 500 body is empty
// // @router / [post]
// func (c *PositionTagController) Post() {
// 	defer c.HandlePanic()
// 	var positionTag models.PositionTag
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &positionTag)
// 	if err != nil {
// 		c.unmarshalErrorHandle(err)
// 		c.ServeJSON()
// 		return
// 	}
// 	c.postHandle(models.CreatePositionTag(positionTag))
// 	c.ServeJSON()
// }

// // Get PositionTag
// // @Title Get
// // @Description get PositionTag by PositionTagID
// // @Param	PositionTagID		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.PositionTag
// // @Failure 500 :PositionTagID is empty
// // @router /:PositionTagID [get]
// func (c *PositionTagController) Get() {
// 	defer c.HandlePanic()
// 	positionTagID, err := c.GetInt64(":PositionTagID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		c.getHandle(models.GetPositionTag(int64(positionTagID)))
// 	}
// 	c.ServeJSON()
// }

// // GetAll ...
// // @Title Get All
// // @Description get PositionTag
// // @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// // @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// // @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// // @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// // @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// // @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// // @Success 200 {object} models.PositionTag
// // @Failure 500
// // @router / [get]
// func (c *PositionTagController) GetAll() {
// 	defer c.HandlePanic()

// 	var limit, offset int64
// 	c.Ctx.Input.Bind(&limit, "limit")
// 	c.Ctx.Input.Bind(&offset, "offset")

// 	c.getHandle(models.GetAllPositionTags(limit, offset))
// 	c.ServeJSON()
// }

// // Put ...
// // @Title Put
// // @Description update the PositionTag
// // @Param	PositionTagID		path 	string	true		"The PositionTagID you want to update"
// // @Param	body		body 	models.PositionTag	true		"body for PositionTag content"
// // @Success 200 {object} models.PositionTag
// // @Failure 500 :PositionTagID is not int
// // @router /:PositionTagID [put]
// func (c *PositionTagController) Put() {
// 	defer c.HandlePanic()
// 	positionTagID, err := c.GetInt64(":PositionTagID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		var positionTag models.PositionTag
// 		err = json.Unmarshal(c.Ctx.Input.RequestBody, &positionTag)
// 		if err != nil {
// 			c.unmarshalErrorHandle(err)
// 		} else {
// 			c.putHandle(nil, models.UpdatePositionTag(int64(positionTagID), &positionTag))
// 		}
// 	}
// 	c.ServeJSON()
// }

// //Delete PositionTag
// // @Title Delete
// // @Description delete the PositionTag
// // @Param  PositionTagID        path    string  true        "The PositionTagID you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 PositionTagID is empty
// // @router /:PositionTagID [delete]
// func (c *PositionTagController) Delete() {
// 	defer c.HandlePanic()
// 	positionTagID, err := c.GetInt64(":PositionTagID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		c.deleteHandle(nil, models.DeletePositionTag(int64(positionTagID)))
// 	}
// 	c.ServeJSON()
// }
