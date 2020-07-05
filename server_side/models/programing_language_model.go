package models

// ProgramingLanguage ...
type ProgramingLanguage struct {
	Model
	ProgramingLanguageName string `gorm:"" json:"programingLanguageName"`
}

// CreateProgramingLanguage ...
func CreateProgramingLanguage(programingLanguage ProgramingLanguage) (ProgramingLanguageID int64, err error) {
	err = db.Create(&programingLanguage).Error
	return programingLanguage.ID, err
}

// GetProgramingLanguage ...
func GetProgramingLanguage(ProgramingLanguageID int64) (programingLanguage ProgramingLanguage, err error) {
	err = db.Set("gorm:auto_preload", true).First(&programingLanguage, ProgramingLanguageID).Error
	return programingLanguage, err
}

// GetAllProgramingLanguages ...
func GetAllProgramingLanguages(limit int64, offset int64) (ml []*ProgramingLanguage, err error) {
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

// UpdateProgramingLanguage ...
func UpdateProgramingLanguage(ProgramingLanguageID int64, programingLanguage *ProgramingLanguage) (err error) {
	err = db.Model(&ProgramingLanguage{Model: Model{ID: ProgramingLanguageID}}).Update(programingLanguage).Error
	return err
}

// DeleteProgramingLanguage ...
func DeleteProgramingLanguage(ProgramingLanguageID int64) (err error) {
	err = db.Delete(&ProgramingLanguage{Model: Model{ID: ProgramingLanguageID}}).Error
	return err
}
