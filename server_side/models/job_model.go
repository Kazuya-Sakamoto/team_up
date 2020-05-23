package models

import "time"

// Job ...
type Job struct {
	Model
	JobTitle           string    `gorm:"" json:"jobTitle"`           // タイトル
	RecruitmentNumbers int64     `gorm:"" json:"recruitmentNumbers"` // 募集人数
	DevStartDate       time.Time `gorm:"" json:"devStartDate"`       // 開発予定開始日
	DevEndDate         time.Time `gorm:"" json:"DevEndDate"`         // 開発予定終了日
	// JobStatusID         int64     `gorm:"" json:"jobStatusId"`         // 募集ステータス
	JobDescription    string    `gorm:"" json:"jobDescription"`    // 募集内容
	PublicationPeriod time.Time `gorm:"" json:"publicationPeriod"` // 掲載期限
	// CommunicationToolID int64     `gorm:"" json:"communicationToolId"` // コミュニケーションツール
	UseMenter bool `gorm:"" json:"useMenter"` // メンター使用の要否
}

// CreateJob ...
func CreateJob(job Job) (JobID int64, err error) {
	err = db.Create(&job).Error
	return job.ID, err
}

// GetJob ...
func GetJob(JobID int64) (job Job, err error) {
	err = db.Set("gorm:auto_preload", true).First(&job, JobID).Error
	return job, err
}

// GetAllJobs ...
func GetAllJobs(limit int64, offset int64) (ml []*Job, err error) {
	tx := db.Set("gorm:auto_preload", true).Begin()

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

// UpdateJob ...
func UpdateJob(JobID int64, job *Job) (err error) {
	err = db.Model(&Job{Model: Model{ID: JobID}}).Update(job).Error
	return err
}

// DeleteJob ...
func DeleteJob(JobID int64) (err error) {
	err = db.Delete(&Job{Model: Model{ID: JobID}}).Error
	return err
}
