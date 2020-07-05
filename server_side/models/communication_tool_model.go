package models

// CommunicationTool ...
type CommunicationTool struct {
	Model
	ToolName string `gorm:"" json:"toolName"`
}

// CreateCommunicationTool ...
func CreateCommunicationTool(communicationTool CommunicationTool) (CommunicationToolID int64, err error) {
	err = db.Create(&communicationTool).Error
	return communicationTool.ID, err
}

// GetCommunicationTool ...
func GetCommunicationTool(CommunicationToolID int64) (communicationTool CommunicationTool, err error) {
	err = db.Set("gorm:auto_preload", true).First(&communicationTool, CommunicationToolID).Error
	return communicationTool, err
}

// GetAllCommunicationTools ...
func GetAllCommunicationTools(limit int64, offset int64) (ml []*CommunicationTool, err error) {
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

// UpdateCommunicationTool ...
func UpdateCommunicationTool(CommunicationToolID int64, communicationTool *CommunicationTool) (err error) {
	err = db.Model(&CommunicationTool{Model: Model{ID: CommunicationToolID}}).Update(communicationTool).Error
	return err
}

// DeleteCommunicationTool ...
func DeleteCommunicationTool(CommunicationToolID int64) (err error) {
	err = db.Delete(&CommunicationTool{Model: Model{ID: CommunicationToolID}}).Error
	return err
}
