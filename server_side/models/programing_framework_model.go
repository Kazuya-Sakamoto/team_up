package models

// ProgramingFramework ...
type ProgramingFramework struct {
	Model
	ProgramingFrameworkName string `gorm:"" json:"programingFrameworkName"`
}

// CreateProgramingFramework ...
func CreateProgramingFramework(programingFramework ProgramingFramework) (ProgramingFrameworkID int64, err error) {
	err = db.Create(&programingFramework).Error
	return programingFramework.ID, err
}

// GetProgramingFramework ...
func GetProgramingFramework(ProgramingFrameworkID int64) (programingFramework ProgramingFramework, err error) {
	err = db.Set("gorm:auto_preload", true).First(&programingFramework, ProgramingFrameworkID).Error
	return programingFramework, err
}

// GetAllProgramingFrameworks ...
func GetAllProgramingFrameworks(limit int64, offset int64) (ml []*ProgramingFramework, err error) {
	tx := db.Set("gorm:auto_preload", true).Begin()

	if limit != 0 {
		tx = tx.Limit(limit)
	} else {
		var count int64
		db.Model(&ml).Count(&count)
		tx = tx.Limit(count)
	}

	err = tx.Offset(offset).Find(&ml).Commit().Error

	return ml, err
}

// UpdateProgramingFramework ...
func UpdateProgramingFramework(ProgramingFrameworkID int64, programingFramework *ProgramingFramework) (err error) {
	err = db.Model(&ProgramingFramework{Model: Model{ID: ProgramingFrameworkID}}).Update(programingFramework).Error
	return err
}

// DeleteProgramingFramework ...
func DeleteProgramingFramework(ProgramingFrameworkID int64) (err error) {
	err = db.Delete(&ProgramingFramework{Model: Model{ID: ProgramingFrameworkID}}).Error
	return err
}
