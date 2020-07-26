package controllers

import (
	"app/server_side/models"
	"app/server_side/services"
	"encoding/json"
)

// ApplyJobController ...
type ApplyJobController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *ApplyJobController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ApplyJob
// @Param	body		body 	models.ApplyJob	true		"body for ApplyJob content"
// @Success 201 {int} models.ApplyJob
// @Failure 500 body is empty
// @router / [post]
func (c *ApplyJobController) Post() {
	defer c.HandlePanic()
	var applyJob models.ApplyJob
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &applyJob)
	if err != nil {
		c.unmarshalErrorHandle(err)
		c.ServeJSON()
		return
	}
	c.postHandle(services.PostApplyJob(applyJob))
	c.ServeJSON()
}

// Get ApplyJob
// @Title Get
// @Description get ApplyJob by ApplyJobID
// @Param	ApplyJobID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ApplyJob
// @Failure 500 :ApplyJobID is empty
// @router /:ApplyJobID [get]
func (c *ApplyJobController) Get() {
	defer c.HandlePanic()
	applyJobID, err := c.GetInt64(":ApplyJobID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetApplyJob(int64(applyJobID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ApplyJob
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ApplyJob
// @Failure 500
// @router / [get]
func (c *ApplyJobController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	var userID, jobID, applyStatusID int64
	c.Ctx.Input.Bind(&userID, "user_id")
	c.Ctx.Input.Bind(&jobID, "job_id")
	c.Ctx.Input.Bind(&applyStatusID, "apply_status_id")

	c.getHandle(models.GetAllApplyJobs(limit, offset, userID, jobID, applyStatusID))
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ApplyJob
// @Param	ApplyJobID		path 	string	true		"The ApplyJobID you want to update"
// @Param	body		body 	models.ApplyJob	true		"body for ApplyJob content"
// @Success 200 {object} models.ApplyJob
// @Failure 500 :ApplyJobID is not int
// @router / [put]
func (c *ApplyJobController) Put() {
	defer c.HandlePanic()
	var applyJob models.ApplyJob
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &applyJob)
	if err != nil {
		c.unmarshalErrorHandle(err)
	} else {
		c.putHandle(nil, services.PutApplyJobWithUserIDAndJobID(applyJob))
	}

	c.ServeJSON()
}

//Delete ApplyJob
// @Title Delete
// @Description delete the ApplyJob
// @Param  ApplyJobID        path    string  true        "The ApplyJobID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 ApplyJobID is empty
// @router / [delete]
func (c *ApplyJobController) Delete() {
	defer c.HandlePanic()
	var applyJob models.ApplyJob
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &applyJob)
	if err != nil {
		c.unmarshalErrorHandle(err)
	} else {
		c.deleteHandle(nil, services.DeleteApplyJobWithUserIDAndJobID(applyJob))
	}

	c.ServeJSON()
}
