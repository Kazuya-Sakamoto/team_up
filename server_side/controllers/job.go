package controllers

import (
	"app/server_side/models"
	"encoding/json"
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

//Post Job
// @Title Post
// @Description create Job
// @Param  body        body    models.Job   true        "body for post content"
// @Success 201 {int} models.Job.Id
// @Failure 403 body is empty
// @router / [post]
func (c *JobController) Post() {
	defer c.HandlePanic()
	var job models.Job
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &job)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	jobID, err := models.CreateJob(job)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		c.Data["json"] = map[string]int64{"jobId": jobID}
		c.Ctx.Output.SetStatus(201)
	}
	c.ServeJSON()
}

//GetAll Jobs
// @Title GetAll
// @Description get all Jobs
// @Success 200 {object} models.Job
// @router / [get]
func (c *JobController) GetAll() {
	defer c.HandlePanic()
	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")
	jobs, err := models.GetAllJobs(limit, offset)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		c.Data["json"] = jobs
	}

	c.ServeJSON()
}

//Get Job
// @Title Get
// @Description get Job by JobID
// @Param  JobID        path    string  true        "The key for staticblock"
// @Success 200 {object} models.Job
// @Failure 403 :JobID is empty
// @router /:JobID [get]
func (c *JobController) Get() {
	defer c.HandlePanic()
	jobID, err := c.GetInt64(":JobID")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		job, err := models.GetJob(int64(jobID))
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(403)
		} else {
			c.Data["json"] = job
		}
	}
	c.ServeJSON()
}

//Put Job
// @Title Update
// @Description update the Job
// @Param  JobID        path    string  true        "The JobID you want to update"
// @Param  body        body    models.Job   true        "body for Job content"
// @Success 200 {object} models.Job
// @Failure 403 :JobID is not int
// @router /:JobID [put]
func (c *JobController) Put() {
	defer c.HandlePanic()
	jobID, err := c.GetInt64(":JobID")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		var job models.Job
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &job)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}
		err = models.UpdateJob(int64(jobID), &job)
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(403)
		} else {
			c.Data["json"] = "update success!"
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
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	} else {
		err = models.DeleteJob(int64(jobID))
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(403)
		} else {
			c.Data["json"] = "delete success!"
		}
	}
	c.ServeJSON()
}
