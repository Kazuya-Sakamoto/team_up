package models

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

// FavoriteJob ...
type FavoriteJob struct {
	Model
	UserID int64 `gorm:"" json:"userId"`
	JobID  int64 `gorm:"" json:"jobId"`
	Job    *Job  `gorm:"" json:"job"`
}

// CreateFavoriteJob ...
func CreateFavoriteJob(tx *gorm.DB, favoriteJob FavoriteJob) (FavoriteJobID int64, err error) {
	err = tx.Create(&favoriteJob).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return favoriteJob.ID, err
}

// FindFavoriteJobWithUserIDAndJobID ...
func FindFavoriteJobWithUserIDAndJobID(tx *gorm.DB, userID int64, jobID int64) (favoriteJob []*FavoriteJob, err error) {
	if userID != 0 && jobID != 0 {
		tx = tx.Where("user_id = ?", userID).Where("job_id = ?", jobID)
	} else {
		logs.Error(err)
		tx.Rollback()
		return favoriteJob, errors.New("userIDもしくはjobIDが不足しています。")
	}

	err = tx.Find(&favoriteJob).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return favoriteJob, err
}

// GetFavoriteJob ...
func GetFavoriteJob(FavoriteJobID int64) (favoriteJob FavoriteJob, err error) {
	err = db.Set("gorm:auto_preload", true).First(&favoriteJob, FavoriteJobID).Error
	return favoriteJob, err
}

// GetAllFavoriteJobs ...
func GetAllFavoriteJobs(limit int64, offset int64, userID int64) (ml []*FavoriteJob, err error) {
	tx := db.Set("gorm:auto_preload", true).Begin()

	if userID != 0 {
		tx = tx.Where("user_id = ?", userID)
	} else {
		logs.Error(err)
		tx.Rollback()
		return ml, errors.New("userIDが必要です。")
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

// UpdateFavoriteJob ...
func UpdateFavoriteJob(FavoriteJobID int64, favoriteJob *FavoriteJob) (err error) {
	err = db.Model(&FavoriteJob{Model: Model{ID: FavoriteJobID}}).Update(favoriteJob).Error
	return err
}

// DeleteFavoriteJob ...
func DeleteFavoriteJob(tx *gorm.DB, userID int64, jobID int64) (err error) {
	err = tx.Where("user_id = ?", userID).Where("job_id = ?", jobID).Unscoped().Delete(&FavoriteJob{}).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return err
}
