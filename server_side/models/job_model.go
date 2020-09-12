package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

// Job ...
type Job struct {
	Model
	JobTitle           string     `gorm:"" json:"jobTitle"`           // タイトル
	RecruitmentNumbers int64      `gorm:"" json:"recruitmentNumbers"` // 募集人数
	DevStartDate       *time.Time `gorm:"" json:"devStartDate"`       // 開発予定開始日
	DevEndDate         *time.Time `gorm:"" json:"devEndDate"`         // 開発予定終了日
	JobStatusID        int64      `gorm:"" json:"jobStatusId"`        // 募集ステータス
	JobDescription     string     `gorm:"" json:"jobDescription"`     // 募集内容
	PublicationPeriod  *time.Time `gorm:"" json:"publicationPeriod"`  // 掲載期限
	// CommunicationToolID  int64                  `gorm:"" json:"communicationToolId"` // コミュニケーションツールID
	// CommunicationTool    *CommunicationTool     `gorm:"" json:"communicationTool"`   // ミュニケーションツール
	UseMenter            bool                   `gorm:"" json:"useMenter"` // メンター使用の要否
	UserID               int64                  `gorm:"" json:"userId"`    // ユーザーID
	User                 *User                  `gorm:"" json:"user"`      // ユーザー
	ProgramingLanguages  []*ProgramingLanguage  `gorm:"many2many:job_programing_languages; association_autoupdate:false" json:"programingLanguage"`
	ProgramingFrameworks []*ProgramingFramework `gorm:"many2many:job_programing_frameworks; association_autoupdate:false" json:"programingFramework"`
	Skills               []*Skill               `gorm:"many2many:job_skills; association_autoupdate:false" json:"skill"`
	// PositionTags         []*PositionTag         `gorm:"many2many:job_position_tags; association_autoupdate:false" json:"positionTag"`
}

// CreateJob ...
func CreateJob(tx *gorm.DB, job Job) (JobID int64, err error) {
	err = tx.Create(&job).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return job.ID, err
}

// GetJob ...
func GetJob(JobID int64) (job Job, err error) {
	err = db.Set("gorm:auto_preload", true).First(&job, JobID).Error
	return job, err
}

// FindJobWithIDAndUserID ...
func FindJobWithIDAndUserID(tx *gorm.DB, jobID int64, userID int64) (job []*Job, err error) {
	if userID != 0 && jobID != 0 {
		tx = tx.Where("user_id = ?", userID).Where("id = ?", jobID)
	} else {
		logs.Error(err)
		tx.Rollback()
		return job, errors.New("userIDもしくはjobIDが不足しています。")
	}

	err = tx.Find(&job).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return
	}
	return job, err
}

// GetAllJobs ...
func GetAllJobs(limit int64, offset int64, positionTagID int64, programingLanguageIDs []int64, skillID int64, devStartDate time.Time, keyword string, userID int64) (ml []*Job, err error) {
	tx := db.Set("gorm:auto_preload", true).Begin()

	if positionTagID != 0 {
		tx = tx.
			Where("id in (?)",
				tx.
					Table("job_position_tags").
					Select("distinct(job_id)").
					Where("position_tag_id = ?", positionTagID).
					SubQuery())
	}
	if len(programingLanguageIDs) != 0 {
		tx = tx.
			Where("id in (?)",
				tx.
					Table("job_programing_languages").
					Select("distinct(job_id)").
					Where("programing_language_id IN (?)", programingLanguageIDs).
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
	if !devStartDate.UTC().IsZero() {
		tx = tx.Where("dev_start_date >= date(?)", devStartDate)
	}
	if keyword != "" {
		tx = tx.Where("job_title like ?", "%%"+keyword+"%%").Or("job_description like ?", "%%"+keyword+"%%")
	}

	if userID != 0 {
		tx = tx.Where("user_id = ?", userID)
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
