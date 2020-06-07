package models

// ChatMessage ...
type ChatMessage struct {
	Model
	Message string `gorm:"" json:"message"`
	JobID   int64  `gorm:"" json:"jobId"`
	UserID  int64  `gorm:"" json:"userId"`
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
func GetAllChatMessages(limit int64, offset int64) (ml []*ChatMessage, err error) {
	tx := db.Begin()

	if limit != 0 {
		tx = tx.Limit(limit)
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