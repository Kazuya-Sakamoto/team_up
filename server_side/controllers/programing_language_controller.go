package controllers

import (
	"app/server_side/models"
	"encoding/json"
)

// ProgramingLanguageController ...
type ProgramingLanguageController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *ProgramingLanguageController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ProgramingLanguage
// @Param	body		body 	models.ProgramingLanguage	true		"body for ProgramingLanguage content"
// @Success 201 {int} models.ProgramingLanguage
// @Failure 500 body is empty
// @router / [post]
func (c *ProgramingLanguageController) Post() {
	defer c.HandlePanic()
	var programingLanguage models.ProgramingLanguage
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &programingLanguage)
	if err != nil {
		c.unmarshalErrorHandle(err)
		c.ServeJSON()
		return
	}
	c.postHandle(models.CreateProgramingLanguage(programingLanguage))
	c.ServeJSON()
}

// Get ProgramingLanguage
// @Title Get
// @Description get ProgramingLanguage by ProgramingLanguageID
// @Param	ProgramingLanguageID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ProgramingLanguage
// @Failure 500 :ProgramingLanguageID is empty
// @router /:ProgramingLanguageID [get]
func (c *ProgramingLanguageController) Get() {
	defer c.HandlePanic()
	programingLanguageID, err := c.GetInt64(":ProgramingLanguageID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetProgramingLanguage(int64(programingLanguageID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ProgramingLanguage
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ProgramingLanguage
// @Failure 500
// @router / [get]
func (c *ProgramingLanguageController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	c.getHandle(models.GetAllProgramingLanguages(limit, offset))
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ProgramingLanguage
// @Param	ProgramingLanguageID		path 	string	true		"The ProgramingLanguageID you want to update"
// @Param	body		body 	models.ProgramingLanguage	true		"body for ProgramingLanguage content"
// @Success 200 {object} models.ProgramingLanguage
// @Failure 500 :ProgramingLanguageID is not int
// @router /:ProgramingLanguageID [put]
func (c *ProgramingLanguageController) Put() {
	defer c.HandlePanic()
	programingLanguageID, err := c.GetInt64(":ProgramingLanguageID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		var programingLanguage models.ProgramingLanguage
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &programingLanguage)
		if err != nil {
			c.unmarshalErrorHandle(err)
		} else {
			c.putHandle(nil, models.UpdateProgramingLanguage(int64(programingLanguageID), &programingLanguage))
		}
	}
	c.ServeJSON()
}

//Delete ProgramingLanguage
// @Title Delete
// @Description delete the ProgramingLanguage
// @Param  ProgramingLanguageID        path    string  true        "The ProgramingLanguageID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 ProgramingLanguageID is empty
// @router /:ProgramingLanguageID [delete]
func (c *ProgramingLanguageController) Delete() {
	defer c.HandlePanic()
	programingLanguageID, err := c.GetInt64(":ProgramingLanguageID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.deleteHandle(nil, models.DeleteProgramingLanguage(int64(programingLanguageID)))
	}
	c.ServeJSON()
}
