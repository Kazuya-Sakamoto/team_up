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
	if err != nil {
		logs.Error(err)
		return
	}
	if len(currentApplyJob) != 0 {
		tx.Rollback()
		return 0, errors.New("この案件はすでに応募済みです。")
	}
	var currentJob []*models.Job
	currentJob, err = models.FindJobWithIDAndUserID(tx, applyJob.JobID, applyJob.UserID)
	if err != nil {
		logs.Error(err)
		return
	}
	if len(currentJob) != 0 {
		tx.Rollback()
		return 0, errors.New("自分の案件に応募することはできません。")
	}
	applyJobID, err = models.CreateApplyJob(tx, applyJob)
	if err != nil {
		logs.Error(err)
		return
	}
	err = tx.Commit().Error
	return
}

// PutApplyJobWithUserIDAndJobID ...
func PutApplyJobWithUserIDAndJobID(applyJob models.ApplyJob) (err error) {
	tx := db.Begin()
	var currentApplyJobs []*models.ApplyJob
	currentApplyJobs, err = models.FindApplyJobWithUserIDAndJobID(tx, applyJob.UserID, applyJob.JobID)
	if err != nil {
		logs.Error(err)
		return
	}
	if len(currentApplyJobs) != 0 {
		currentApplyJob := currentApplyJobs[0]
		err = models.UpdateApplyJob(tx, currentApplyJob.ID, applyJob)
		if err != nil {
			logs.Error(err)
			return
		}
	} else {
		tx.Rollback()
		return errors.New("対象の案件がありません。")
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
