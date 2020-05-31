package models

import (
	"time"
)

// Job ...
type Job struct {
	Model
	JobTitle             string                 `gorm:"" json:"jobTitle"`            // タイトル
	RecruitmentNumbers   int64                  `gorm:"" json:"recruitmentNumbers"`  // 募集人数
	DevStartDate         time.Time              `gorm:"" json:"devStartDate"`        // 開発予定開始日
	DevEndDate           time.Time              `gorm:"" json:"devEndDate"`          // 開発予定終了日
	JobStatusID          int64                  `gorm:"" json:"jobStatusId"`         // 募集ステータス
	JobDescription       string                 `gorm:"" json:"jobDescription"`      // 募集内容
	PublicationPeriod    time.Time              `gorm:"" json:"publicationPeriod"`   // 掲載期限
	CommunicationToolID  int64                  `gorm:"" json:"communicationToolId"` // コミュニケーションツール
	UseMenter            bool                   `gorm:"" json:"useMenter"`           // メンター使用の要否
	ProgramingLanguages  []*ProgramingLanguage  `gorm:"many2many:job_programing_languages; association_autoupdate:false" json:"programingLanguage"`
	ProgramingFrameworks []*ProgramingFramework `gorm:"many2many:job_programing_frameworks; association_autoupdate:false" json:"programingFramework"`
	Skills               []*Skill               `gorm:"many2many:job_skills; association_autoupdate:false" json:"skill"`
	PositionTags         []*PositionTag         `gorm:"many2many:job_position_tags; association_autoupdate:false" json:"positionTag"`
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
func GetAllJobs(limit int64, offset int64, positionTagID int64, programingLanguageID int64, skillID int64, devStartDate time.Time) (ml []*Job, err error) {
	tx := db.Begin()

	if positionTagID != 0 {
		tx = tx.
			Where("id in (?)",
				tx.
					Table("job_position_tags").
					Select("distinct(job_id)").
					Where("position_tag_id = ?", positionTagID).
					SubQuery())
	}
	if programingLanguageID != 0 {
		tx = tx.
			Where("id in (?)",
				tx.
					Table("job_programing_languages").
					Select("distinct(job_id)").
					Where("programing_language_id = ?", programingLanguageID).
					SubQuery())
	}
	if skillID != 0 {
		tx = tx.
			Where("id in (?)",
				tx.
					Table("job_skills").
					Select("distinct(job_id)").
					Where("skill_id = ?", skillID).
					SubQuery())
	}
	if devStartDate.IsZero() == false {
		tx = tx.Where("dev_start_date >= ?", devStartDate)
	}
	if limit != 0 {
		tx = tx.Limit(limit)
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
