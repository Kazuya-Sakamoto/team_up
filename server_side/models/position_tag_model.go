package models

// // PositionTag ...
// type PositionTag struct {
// 	Model
// 	PositionTagName string `gorm:"" json:"positionTagName"`
// }

// // CreatePositionTag ...
// func CreatePositionTag(positionTag PositionTag) (PositionTagID int64, err error) {
// 	err = db.Create(&positionTag).Error
// 	return positionTag.ID, err
// }

// // GetPositionTag ...
// func GetPositionTag(PositionTagID int64) (positionTag PositionTag, err error) {
// 	err = db.Set("gorm:auto_preload", true).First(&positionTag, PositionTagID).Error
// 	return positionTag, err
// }

// // GetAllPositionTags ...
// func GetAllPositionTags(limit int64, offset int64) (ml []*PositionTag, err error) {
// 	tx := db.Begin()

// 	if limit != 0 {
// 		tx = tx.Limit(limit)
// 	} else {
// 		var count int64
// 		tx.Model(&ml).Count(&count)
// 		tx = tx.Limit(count)
// 	}

// 	err = tx.Offset(offset).Find(&ml).Commit().Error

// 	return ml, err
// }

// // UpdatePositionTag ...
// func UpdatePositionTag(PositionTagID int64, positionTag *PositionTag) (err error) {
// 	err = db.Model(&PositionTag{Model: Model{ID: PositionTagID}}).Update(positionTag).Error
// 	return err
// }

// // DeletePositionTag ...
// func DeletePositionTag(PositionTagID int64) (err error) {
// 	err = db.Delete(&PositionTag{Model: Model{ID: PositionTagID}}).Error
// 	return err
// }
