package models

// JobStatus ...
type JobStatus struct {
	Model
	JobStatus string `gorm:"" json:"jobStatus"`
}

// CreateJobStatus ...
func CreateJobStatus(jobStatus JobStatus) (JobStatusID int64, err error) {
	err = db.Create(&jobStatus).Error
	return jobStatus.ID, err
}

// GetJobStatus ...
func GetJobStatus(JobStatusID int64) (jobStatus JobStatus, err error) {
	err = db.Set("gorm:auto_preload", true).First(&jobStatus, JobStatusID).Error
	return jobStatus, err
}

// GetAllJobStatuses ...
func GetAllJobStatuses(limit int64, offset int64) (ml []*JobStatus, err error) {
	tx := db.Begin()

	if limit != 0 {
		tx = tx.Limit(limit)
	}

	err = tx.Offset(offset).Find(&ml).Commit().Error

	return ml, err
}

// UpdateJobStatus ...
func UpdateJobStatus(JobStatusID int64, jobStatus *JobStatus) (err error) {
	err = db.Model(&JobStatus{Model: Model{ID: JobStatusID}}).Update(jobStatus).Error
	return err
}

// DeleteJobStatus ...
func DeleteJobStatus(JobStatusID int64) (err error) {
	err = db.Delete(&JobStatus{Model: Model{ID: JobStatusID}}).Error
	return err
}
