package controllers

import (
	"app/server_side/models"
	"encoding/json"
	"time"
)

// JobController ...
type JobController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *JobController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Job
// @Param	body		body 	models.Job	true		"body for Job content"
// @Success 201 {int} models.Job
// @Failure 500 body is empty
// @router / [post]
func (c *JobController) Post() {
	defer c.HandlePanic()
	var job models.Job
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &job)
	if err != nil {
		c.unmarshalErrorHandle(err)
		c.ServeJSON()
		return
	}
	c.postHandle(models.CreateJob(job))
	c.ServeJSON()
}

// Get Job
// @Title Get
// @Description get Job by JobID
// @Param	JobID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Job
// @Failure 500 :JobID is empty
// @router /:JobID [get]
func (c *JobController) Get() {
	defer c.HandlePanic()
	jobID, err := c.GetInt64(":JobID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetJob(int64(jobID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Job
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Job
// @Failure 500
// @router / [get]
func (c *JobController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	var positionTagID, programingLanguageID, skillID int64
	var keyword string
	c.Ctx.Input.Bind(&positionTagID, "position_tag_id")
	c.Ctx.Input.Bind(&programingLanguageID, "programing_language_id")
	c.Ctx.Input.Bind(&skillID, "skill_id")
	c.Ctx.Input.Bind(&keyword, "keyword")

	var devStartDateStr string
	var devStartDate time.Time
	c.Ctx.Input.Bind(&devStartDateStr, "dev_start_date")
	devStartDate, _ = time.Parse("2006-01-02", devStartDateStr)

	var userID int64
	c.Ctx.Input.Bind(&userID, "user_id")

	c.getHandle(models.GetAllJobs(limit, offset, positionTagID, programingLanguageID, skillID, devStartDate, keyword, userID))
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Job
// @Param	JobID		path 	string	true		"The JobID you want to update"
// @Param	body		body 	models.Job	true		"body for Job content"
// @Success 200 {object} models.Job
// @Failure 500 :JobID is not int
// @router /:JobID [put]
func (c *JobController) Put() {
	defer c.HandlePanic()
	jobID, err := c.GetInt64(":JobID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		var job models.Job
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &job)
		if err != nil {
			c.unmarshalErrorHandle(err)
		} else {
			c.putHandle(nil, models.UpdateJob(int64(jobID), &job))
		}
	}
	c.ServeJSON()
}

//Delete Job
// @Title Delete
// @Description delete the Job
// @Param  JobID        path    string  true        "The JobID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 JobID is empty
// @router /:JobID [delete]
func (c *JobController) Delete() {
	defer c.HandlePanic()
	jobID, err := c.GetInt64(":JobID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.deleteHandle(nil, models.DeleteJob(int64(jobID)))
	}
	c.ServeJSON()
}
