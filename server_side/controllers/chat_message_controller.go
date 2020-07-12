package controllers

import (
	"app/server_side/models"
	"encoding/json"
)

// ChatMessageController ...
type ChatMessageController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *ChatMessageController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ChatMessage
// @Param	body		body 	models.ChatMessage	true		"body for ChatMessage content"
// @Success 201 {int} models.ChatMessage
// @Failure 500 body is empty
// @router / [post]
func (c *ChatMessageController) Post() {
	defer c.HandlePanic()
	var chatMessage models.ChatMessage
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &chatMessage)
	if err != nil {
		c.unmarshalErrorHandle(err)
		c.ServeJSON()
		return
	}
	c.postHandle(models.CreateChatMessage(chatMessage))
	c.ServeJSON()
}

// Get ChatMessage
// @Title Get
// @Description get ChatMessage by ChatMessageID
// @Param	ChatMessageID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ChatMessage
// @Failure 500 :ChatMessageID is empty
// @router /:ChatMessageID [get]
func (c *ChatMessageController) Get() {
	defer c.HandlePanic()
	chatMessageID, err := c.GetInt64(":ChatMessageID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetChatMessage(int64(chatMessageID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ChatMessage
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ChatMessage
// @Failure 500
// @router / [get]
func (c *ChatMessageController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	var jobID int64
	c.Ctx.Input.Bind(&jobID, "job_id")

	c.getHandle(models.GetAllChatMessages(limit, offset, jobID))
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ChatMessage
// @Param	ChatMessageID		path 	string	true		"The ChatMessageID you want to update"
// @Param	body		body 	models.ChatMessage	true		"body for ChatMessage content"
// @Success 200 {object} models.ChatMessage
// @Failure 500 :ChatMessageID is not int
// @router /:ChatMessageID [put]
func (c *ChatMessageController) Put() {
	defer c.HandlePanic()
	chatMessageID, err := c.GetInt64(":ChatMessageID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		var chatMessage models.ChatMessage
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &chatMessage)
		if err != nil {
			c.unmarshalErrorHandle(err)
		} else {
			c.putHandle(nil, models.UpdateChatMessage(int64(chatMessageID), &chatMessage))
		}
	}
	c.ServeJSON()
}

//Delete ChatMessage
// @Title Delete
// @Description delete the ChatMessage
// @Param  ChatMessageID        path    string  true        "The ChatMessageID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 ChatMessageID is empty
// @router /:ChatMessageID [delete]
func (c *ChatMessageController) Delete() {
	defer c.HandlePanic()
	chatMessageID, err := c.GetInt64(":ChatMessageID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.deleteHandle(nil, models.DeleteChatMessage(int64(chatMessageID)))
	}
	c.ServeJSON()
}
