package models

import "github.com/jinzhu/gorm"

// Skill ...
type Skill struct {
	Model
	SkillName string `gorm:"" json:"skillName"`
}

// CreateSkill ...
func CreateSkill(skill Skill) (SkillID int64, err error) {
	err = db.Create(&skill).Error
	return skill.ID, err
}

// GetSkill ...
func GetSkill(SkillID int64) (skill Skill, err error) {
	err = db.Set("gorm:auto_preload", true).First(&skill, SkillID).Error
	return skill, err
}

// GetAllSkills ...
func GetAllSkills(limit int64, offset int64) (ml []*Skill, err error) {
	tx := db.Begin()

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

// UpdateSkill ...
func UpdateSkill(SkillID int64, skill *Skill) (err error) {
	err = db.Model(&Skill{Model: Model{ID: SkillID}}).Update(skill).Error
	return err
}

// DeleteSkill ...
func DeleteSkill(SkillID int64) (err error) {
	err = db.Delete(&Skill{Model: Model{ID: SkillID}}).Error
	return err
}

// CountCreateSkills ...
func CountCreateSkills(tx *gorm.DB, SKs []*Skill) (count int) {
	count = len(SKs)
	return count
}
