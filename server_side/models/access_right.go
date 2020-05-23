package models

// // AccessRight ...
// type AccessRight struct {
// 	Model
// 	AccessRightName string  `gorm:"" json:"accessRightName"`
// 	Roles           []*Role `gorm:"many2many:role_access_rights" json:"roles"`
// }

// func init() {}

// // CreateAccessRight ...
// func CreateAccessRight(accessRight AccessRight) (accessRightID int64, err error) {
// 	err = db.Create(&accessRight).Error
// 	return
// }

// // GetAccessRight ...
// func GetAccessRight(accessRightID int64) (accessRight AccessRight, err error) {
// 	err = db.First(&accessRight, accessRightID).Error
// 	return
// }

// // GetAllAccessRights ...
// func GetAllAccessRights(limit int64, offset int64) (ml []*AccessRight, err error) {
// 	tx := db.Begin()

// 	if limit != 0 {
// 		tx = tx.Limit(limit)
// 	} else {
// 		var count int64
// 		db.Model(&ml).Count(&count)
// 		tx = tx.Limit(count)
// 	}

// 	err = tx.Offset(offset).Find(&ml).Commit().Error
// 	return
// }

// // UpdateAccessRight ...
// func UpdateAccessRight(AccessRightID int64, accessRight *AccessRight) (err error) {
// 	err = db.Model(&AccessRight{Model: Model{ID: AccessRightID}}).Update(accessRight).Error
// 	return
// }

// // DeleteAccessRight ...
// func DeleteAccessRight(AccessRightID int64) (err error) {
// 	err = db.Delete(&AccessRight{Model: Model{ID: AccessRightID}}).Error
// 	return
// }
