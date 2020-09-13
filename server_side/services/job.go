package services

import (
	"app/server_side/models"
	"errors"

	"github.com/astaxie/beego/logs"
)

// PostJob ...
func PostJob(job models.Job) (JobID int64, err error) {
	tx := db.Begin()

	countPLs := models.CountCreateProgramingLanguages(tx, job.ProgramingLanguages)
	if countPLs > 5 {
		logs.Error(err)
		tx.Rollback()
		return JobID, errors.New("プログラミング言語は5個までしか登録できません。")
	}

	countPFs := models.CountCreateProgramingFrameworks(tx, job.ProgramingFrameworks)
	if countPFs > 5 {
		logs.Error(err)
		tx.Rollback()
		return JobID, errors.New("フレームワークは5個までしか登録できません。")
	}

	countSKs := models.CountCreateSkills(tx, job.Skills)
	if countSKs > 5 {
		logs.Error(err)
		tx.Rollback()
		return JobID, errors.New("その他スキルは5個までしか登録できません。")
	}

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
