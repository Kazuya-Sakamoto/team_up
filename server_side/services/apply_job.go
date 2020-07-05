package services

import (
	"app/server_side/models"
	"errors"

	"github.com/astaxie/beego/logs"
)

// PostApplyJob ...
func PostApplyJob(applyJob models.ApplyJob) (applyJobID int64, err error) {
	tx := db.Begin()
	var currentApplyJob []*models.ApplyJob
	currentApplyJob, err = models.FindApplyJobWithUserIDAndJobID(tx, applyJob.UserID, applyJob.JobID)
	if len(currentApplyJob) != 0 {
		tx.Rollback()
		return 0, errors.New("この案件はすでに応募済みです。")
	}
	applyJobID, err = models.CreateApplyJob(tx, applyJob)
	if err != nil {
		logs.Error(err)
		return
	}
	err = tx.Commit().Error
	return
}

// DeleteApplyJobWithUserIDAndJobID ...
func DeleteApplyJobWithUserIDAndJobID(applyJob models.ApplyJob) (err error) {
	tx := db.Begin()
	var currentApplyJob []*models.ApplyJob
	currentApplyJob, err = models.FindApplyJobWithUserIDAndJobID(tx, applyJob.UserID, applyJob.JobID)
	if err != nil {
		logs.Error(err)
		return
	}
	if len(currentApplyJob) != 0 {
		err = models.DeleteApplyJob(tx, applyJob.UserID, applyJob.JobID)
		if err != nil {
			logs.Error(err)
			return
		}
	} else {
		tx.Rollback()
		return errors.New("この案件はすでに応募取り消し済みです。")
	}

	err = tx.Commit().Error
	return
}
