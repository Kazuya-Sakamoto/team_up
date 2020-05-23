package controllers

import (
	"app/server_side/models"
	"encoding/json"
)

// ProgramingFrameworkController ...
type ProgramingFrameworkController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *ProgramingFrameworkController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ProgramingFramework
// @Param	body		body 	models.ProgramingFramework	true		"body for ProgramingFramework content"
// @Success 201 {int} models.ProgramingFramework
// @Failure 500 body is empty
// @router / [post]
func (c *ProgramingFrameworkController) Post() {
	defer c.HandlePanic()
	var programingFramework models.ProgramingFramework
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &programingFramework)
	if err != nil {
		c.unmarshalErrorHandle(err)
		c.ServeJSON()
		return
	}
	c.postHandle(models.CreateProgramingFramework(programingFramework))
	c.ServeJSON()
}

// Get ProgramingFramework
// @Title Get
// @Description get ProgramingFramework by ProgramingFrameworkID
// @Param	ProgramingFrameworkID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ProgramingFramework
// @Failure 500 :ProgramingFrameworkID is empty
// @router /:ProgramingFrameworkID [get]
func (c *ProgramingFrameworkController) Get() {
	defer c.HandlePanic()
	programingFrameworkID, err := c.GetInt64(":ProgramingFrameworkID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetProgramingFramework(int64(programingFrameworkID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ProgramingFramework
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ProgramingFramework
// @Failure 500
// @router / [get]
func (c *ProgramingFrameworkController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	c.getHandle(models.GetAllProgramingFrameworks(limit, offset))
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ProgramingFramework
// @Param	ProgramingFrameworkID		path 	string	true		"The ProgramingFrameworkID you want to update"
// @Param	body		body 	models.ProgramingFramework	true		"body for ProgramingFramework content"
// @Success 200 {object} models.ProgramingFramework
// @Failure 500 :ProgramingFrameworkID is not int
// @router /:ProgramingFrameworkID [put]
func (c *ProgramingFrameworkController) Put() {
	defer c.HandlePanic()
	programingFrameworkID, err := c.GetInt64(":ProgramingFrameworkID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		var programingFramework models.ProgramingFramework
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &programingFramework)
		if err != nil {
			c.unmarshalErrorHandle(err)
		} else {
			c.putHandle(nil, models.UpdateProgramingFramework(int64(programingFrameworkID), &programingFramework))
		}
	}
	c.ServeJSON()
}

//Delete ProgramingFramework
// @Title Delete
// @Description delete the ProgramingFramework
// @Param  ProgramingFrameworkID        path    string  true        "The ProgramingFrameworkID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 ProgramingFrameworkID is empty
// @router /:ProgramingFrameworkID [delete]
func (c *ProgramingFrameworkController) Delete() {
	defer c.HandlePanic()
	programingFrameworkID, err := c.GetInt64(":ProgramingFrameworkID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.deleteHandle(nil, models.DeleteProgramingFramework(int64(programingFrameworkID)))
	}
	c.ServeJSON()
}
