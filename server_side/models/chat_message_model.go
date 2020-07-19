package models

import (
	"errors"

	"github.com/astaxie/beego/logs"
)

// ChatMessage ...
type ChatMessage struct {
	Model
	Message string `gorm:"" json:"message"`
	JobID   int64  `gorm:"" json:"jobId"`
	Job     *Job   `gorm:"" json:"job"`
	UserID  int64  `gorm:"" json:"userId"`
	User    *User  `gorm:"" json:"user"`
}

// CreateChatMessage ...
func CreateChatMessage(chatMessage ChatMessage) (ChatMessageID int64, err error) {
	err = db.Create(&chatMessage).Error
	return chatMessage.ID, err
}

// GetChatMessage ...
func GetChatMessage(ChatMessageID int64) (chatMessage ChatMessage, err error) {
	err = db.Set("gorm:auto_preload", true).First(&chatMessage, ChatMessageID).Error
	return chatMessage, err
}

// GetAllChatMessages ...
func GetAllChatMessages(limit int64, offset int64, jobID int64) (ml []*ChatMessage, err error) {
	tx := db.Set("gorm:auto_preload", true).Begin()

	if jobID != 0 {
		tx = tx.Where("job_id = ?", jobID)
	} else {
		logs.Error(err)
		tx.Rollback()
		return ml, errors.New("jobIDが必要です。")
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

// UpdateChatMessage ...
func UpdateChatMessage(ChatMessageID int64, chatMessage *ChatMessage) (err error) {
	err = db.Model(&ChatMessage{Model: Model{ID: ChatMessageID}}).Update(chatMessage).Error
	return err
}

// DeleteChatMessage ...
func DeleteChatMessage(ChatMessageID int64) (err error) {
	err = db.Delete(&ChatMessage{Model: Model{ID: ChatMessageID}}).Error
	return err
}
