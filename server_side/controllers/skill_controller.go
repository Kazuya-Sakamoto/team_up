package controllers

import (
	"app/server_side/models"
	"encoding/json"
)

// SkillController ...
type SkillController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *SkillController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Skill
// @Param	body		body 	models.Skill	true		"body for Skill content"
// @Success 201 {int} models.Skill
// @Failure 500 body is empty
// @router / [post]
func (c *SkillController) Post() {
	defer c.HandlePanic()
	var skill models.Skill
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &skill)
	if err != nil {
		c.unmarshalErrorHandle(err)
		c.ServeJSON()
		return
	}
	c.postHandle(models.CreateSkill(skill))
	c.ServeJSON()
}

// Get Skill
// @Title Get
// @Description get Skill by SkillID
// @Param	SkillID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Skill
// @Failure 500 :SkillID is empty
// @router /:SkillID [get]
func (c *SkillController) Get() {
	defer c.HandlePanic()
	skillID, err := c.GetInt64(":SkillID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetSkill(int64(skillID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Skill
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Skill
// @Failure 500
// @router / [get]
func (c *SkillController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	c.getHandle(models.GetAllSkills(limit, offset))
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Skill
// @Param	SkillID		path 	string	true		"The SkillID you want to update"
// @Param	body		body 	models.Skill	true		"body for Skill content"
// @Success 200 {object} models.Skill
// @Failure 500 :SkillID is not int
// @router /:SkillID [put]
func (c *SkillController) Put() {
	defer c.HandlePanic()
	skillID, err := c.GetInt64(":SkillID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		var skill models.Skill
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &skill)
		if err != nil {
			c.unmarshalErrorHandle(err)
		} else {
			c.putHandle(nil, models.UpdateSkill(int64(skillID), &skill))
		}
	}
	c.ServeJSON()
}

//Delete Skill
// @Title Delete
// @Description delete the Skill
// @Param  SkillID        path    string  true        "The SkillID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 SkillID is empty
// @router /:SkillID [delete]
func (c *SkillController) Delete() {
	defer c.HandlePanic()
	skillID, err := c.GetInt64(":SkillID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.deleteHandle(nil, models.DeleteSkill(int64(skillID)))
	}
	c.ServeJSON()
}
