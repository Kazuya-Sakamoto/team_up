package controllers

// import (
// 	"app/server_side/models"
// 	"encoding/json"
// )

// // CommunicationToolController ...
// type CommunicationToolController struct {
// 	RequiredLoginController
// }

// // URLMapping ...
// func (c *CommunicationToolController) URLMapping() {
// 	c.Mapping("Post", c.Post)
// 	c.Mapping("Get", c.Get)
// 	c.Mapping("GetAll", c.GetAll)
// 	c.Mapping("Put", c.Put)
// 	c.Mapping("Delete", c.Delete)
// }

// // Post ...
// // @Title Post
// // @Description create CommunicationTool
// // @Param	body		body 	models.CommunicationTool	true		"body for CommunicationTool content"
// // @Success 201 {int} models.CommunicationTool
// // @Failure 500 body is empty
// // @router / [post]
// func (c *CommunicationToolController) Post() {
// 	defer c.HandlePanic()
// 	var communicationTool models.CommunicationTool
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &communicationTool)
// 	if err != nil {
// 		c.unmarshalErrorHandle(err)
// 		c.ServeJSON()
// 		return
// 	}
// 	c.postHandle(models.CreateCommunicationTool(communicationTool))
// 	c.ServeJSON()
// }

// // Get CommunicationTool
// // @Title Get
// // @Description get CommunicationTool by CommunicationToolID
// // @Param	CommunicationToolID		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.CommunicationTool
// // @Failure 500 :CommunicationToolID is empty
// // @router /:CommunicationToolID [get]
// func (c *CommunicationToolController) Get() {
// 	defer c.HandlePanic()
// 	communicationToolID, err := c.GetInt64(":CommunicationToolID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		c.getHandle(models.GetCommunicationTool(int64(communicationToolID)))
// 	}
// 	c.ServeJSON()
// }

// // GetAll ...
// // @Title Get All
// // @Description get CommunicationTool
// // @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// // @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// // @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// // @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// // @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// // @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// // @Success 200 {object} models.CommunicationTool
// // @Failure 500
// // @router / [get]
// func (c *CommunicationToolController) GetAll() {
// 	defer c.HandlePanic()

// 	var limit, offset int64
// 	c.Ctx.Input.Bind(&limit, "limit")
// 	c.Ctx.Input.Bind(&offset, "offset")

// 	c.getHandle(models.GetAllCommunicationTools(limit, offset))
// 	c.ServeJSON()
// }

// // Put ...
// // @Title Put
// // @Description update the CommunicationTool
// // @Param	CommunicationToolID		path 	string	true		"The CommunicationToolID you want to update"
// // @Param	body		body 	models.CommunicationTool	true		"body for CommunicationTool content"
// // @Success 200 {object} models.CommunicationTool
// // @Failure 500 :CommunicationToolID is not int
// // @router /:CommunicationToolID [put]
// func (c *CommunicationToolController) Put() {
// 	defer c.HandlePanic()
// 	communicationToolID, err := c.GetInt64(":CommunicationToolID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		var communicationTool models.CommunicationTool
// 		err = json.Unmarshal(c.Ctx.Input.RequestBody, &communicationTool)
// 		if err != nil {
// 			c.unmarshalErrorHandle(err)
// 		} else {
// 			c.putHandle(nil, models.UpdateCommunicationTool(int64(communicationToolID), &communicationTool))
// 		}
// 	}
// 	c.ServeJSON()
// }

// //Delete CommunicationTool
// // @Title Delete
// // @Description delete the CommunicationTool
// // @Param  CommunicationToolID        path    string  true        "The CommunicationToolID you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 CommunicationToolID is empty
// // @router /:CommunicationToolID [delete]
// func (c *CommunicationToolController) Delete() {
// 	defer c.HandlePanic()
// 	communicationToolID, err := c.GetInt64(":CommunicationToolID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		c.deleteHandle(nil, models.DeleteCommunicationTool(int64(communicationToolID)))
// 	}
// 	c.ServeJSON()
// }
