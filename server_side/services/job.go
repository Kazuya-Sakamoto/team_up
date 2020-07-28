package services

import (
	"app/server_side/models"

	"github.com/astaxie/beego/logs"
)

// PostJob ...
func PostJob(job models.Job) (JobID int64, err error) {
	tx := db.Begin()

	JobID, err = models.CreateJob(tx, job)
	if err != nil {
		logs.Error(err)
		return
	}
	applyJob := models.ApplyJob{
		UserID:        job.UserID,
		JobID:         JobID,
		ApplyStatusID: 2,
	}
	_, err = models.CreateApplyJob(tx, applyJob)
	if err != nil {
		logs.Error(err)
		return
	}
	err = tx.Commit().Error
	return
}
