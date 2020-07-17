package models

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

// ApplyJob ...
type ApplyJob struct {
	Model
	UserID        int64        `gorm:"" json:"userId"`
	User          *User        `gorm:"" json:"user"`
	JobID         int64        `gorm:"" json:"jobId"`
	Job           *Job         `gorm:"" json:"job"`
	ApplyStatusID int64        `gorm:"" json:"applyStatusId"`
	ApplyStatus   *ApplyStatus `gorm:"" json:"applyStatus"`
}

// ApplyStatus ...
type ApplyStatus struct {
	Model
	StatusName string `gorm:"" json:"statusName"`
}

// CreateApplyJob ...
func CreateApplyJob(tx *gorm.DB, applyJob ApplyJob) (ApplyJobID int64, err error) {
	err = tx.Create(&applyJob).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return applyJob.ID, err
}

// FindApplyJobWithUserIDAndJobID ...
func FindApplyJobWithUserIDAndJobID(tx *gorm.DB, userID int64, jobID int64) (applyJob []*ApplyJob, err error) {
	if userID != 0 && jobID != 0 {
		tx = tx.Where("user_id = ?", userID).Where("job_id = ?", jobID)
	} else {
		logs.Error(err)
		tx.Rollback()
		return applyJob, errors.New("userIDもしくはjobIDが不足しています。")
	}

	err = tx.Find(&applyJob).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return applyJob, err
}

// GetApplyJob ...
func GetApplyJob(ApplyJobID int64) (applyJob ApplyJob, err error) {
	err = db.Set("gorm:auto_preload", true).First(&applyJob, ApplyJobID).Error
	return applyJob, err
}

// GetAllApplyJobs ...
func GetAllApplyJobs(limit int64, offset int64, userID int64, jobID int64) (ml []*ApplyJob, err error) {
	tx := db.Set("gorm:auto_preload", true).Begin()

	if userID != 0 {
		tx = tx.Where("user_id = ?", userID)
	}
	if jobID != 0 {
		tx = tx.Where("job_id = ?", jobID)
	}
	if userID == 0 && jobID == 0 {
		logs.Error(err)
		tx.Rollback()
		return ml, errors.New("userIDもしくはjobIDが必要です。")
	}

	if limit != 0 {
		tx = tx.Limit(limit)
	} else {
		var count int64
		tx.Model(&ml).Count(&count)
		tx = tx.Limit(count)
	}

	err = tx.Offset(offset).Find(&ml).Commit().Error

	return ml, err
}

// UpdateApplyJob ...
func UpdateApplyJob(ApplyJobID int64, applyJob *ApplyJob) (err error) {
	err = db.Model(&ApplyJob{Model: Model{ID: ApplyJobID}}).Update(applyJob).Error
	return err
}

// DeleteApplyJob ...
func DeleteApplyJob(tx *gorm.DB, userID int64, jobID int64) (err error) {
	err = tx.Where("user_id = ?", userID).Where("job_id = ?", jobID).Unscoped().Delete(&ApplyJob{}).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return err
}
