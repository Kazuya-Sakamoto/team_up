package models

// IndividualPortfolio ...
type IndividualPortfolio struct {
	Model
	IndividualPortfolioURL      string `gorm:"" json:"individualPortfolioUrl"`
	IndividualPortfolioName     string `gorm:"" json:"individualPortfolioName"`
	IndividualPortfolioOverview string `gorm:"" json:"individualPortfolioOverview"`
	UserID                      int64  `gorm:"" json:"userId"`
}

// CreateIndividualPortfolio ...
func CreateIndividualPortfolio(programingLanguage IndividualPortfolio) (IndividualPortfolioID int64, err error) {
	err = db.Create(&programingLanguage).Error
	return programingLanguage.ID, err
}

// GetIndividualPortfolio ...
func GetIndividualPortfolio(IndividualPortfolioID int64) (programingLanguage IndividualPortfolio, err error) {
	err = db.Set("gorm:auto_preload", true).First(&programingLanguage, IndividualPortfolioID).Error
	return programingLanguage, err
}

// GetAllIndividualPortfolios ...
func GetAllIndividualPortfolios(limit int64, offset int64) (ml []*IndividualPortfolio, err error) {
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

// UpdateIndividualPortfolio ...
func UpdateIndividualPortfolio(IndividualPortfolioID int64, programingLanguage *IndividualPortfolio) (err error) {
	err = db.Model(&IndividualPortfolio{Model: Model{ID: IndividualPortfolioID}}).Update(programingLanguage).Error
	return err
}

// DeleteIndividualPortfolio ...
func DeleteIndividualPortfolio(IndividualPortfolioID int64) (err error) {
	err = db.Delete(&IndividualPortfolio{Model: Model{ID: IndividualPortfolioID}}).Error
	return err
}
