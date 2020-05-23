package controllers

import (
	"app/server_side/models"
	"encoding/json"
)

// JobStatusController ...
type JobStatusController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *JobStatusController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create JobStatus
// @Param	body		body 	models.JobStatus	true		"body for JobStatus content"
// @Success 201 {int} models.JobStatus
// @Failure 500 body is empty
// @router / [post]
func (c *JobStatusController) Post() {
	defer c.HandlePanic()
	var jobStatus models.JobStatus
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &jobStatus)
	if err != nil {
		c.unmarshalErrorHandle(err)
		c.ServeJSON()
		return
	}
	c.postHandle(models.CreateJobStatus(jobStatus))
	c.ServeJSON()
}

// Get JobStatus
// @Title Get
// @Description get JobStatus by JobStatusID
// @Param	JobStatusID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.JobStatus
// @Failure 500 :JobStatusID is empty
// @router /:JobStatusID [get]
func (c *JobStatusController) Get() {
	defer c.HandlePanic()
	jobStatusID, err := c.GetInt64(":JobStatusID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetJobStatus(int64(jobStatusID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get JobStatus
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.JobStatus
// @Failure 500
// @router / [get]
func (c *JobStatusController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	c.getHandle(models.GetAllJobStatuses(limit, offset))
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the JobStatus
// @Param	JobStatusID		path 	string	true		"The JobStatusID you want to update"
// @Param	body		body 	models.JobStatus	true		"body for JobStatus content"
// @Success 200 {object} models.JobStatus
// @Failure 500 :JobStatusID is not int
// @router /:JobStatusID [put]
func (c *JobStatusController) Put() {
	defer c.HandlePanic()
	jobStatusID, err := c.GetInt64(":JobStatusID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		var jobStatus models.JobStatus
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &jobStatus)
		if err != nil {
			c.unmarshalErrorHandle(err)
		} else {
			c.putHandle(nil, models.UpdateJobStatus(int64(jobStatusID), &jobStatus))
		}
	}
	c.ServeJSON()
}

//Delete JobStatus
// @Title Delete
// @Description delete the JobStatus
// @Param  JobStatusID        path    string  true        "The JobStatusID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 JobStatusID is empty
// @router /:JobStatusID [delete]
func (c *JobStatusController) Delete() {
	defer c.HandlePanic()
	jobStatusID, err := c.GetInt64(":JobStatusID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.deleteHandle(nil, models.DeleteJobStatus(int64(jobStatusID)))
	}
	c.ServeJSON()
}
