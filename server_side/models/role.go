package models

// // Role ...
// type Role struct {
// 	Model
// 	RoleName     string         `gorm:"" json:"roleName"`
// 	AccessRights []*AccessRight `gorm:"many2many:role_access_rights;save_association:false" json:"accessRights"`
// }

// func init() {

// }

// // CreateRole ...
// func CreateRole(role Role) (roleID int64, err error) {
// 	err = db.Create(&role).Error
// 	return
// }

// // GetRole ...
// func GetRole(roleID int64) (role Role, err error) {
// 	err = db.Set("gorm:auto_preload", true).First(&role, roleID).Error
// 	return
// }

// // GetAllRoles ...
// func GetAllRoles(limit int64, offset int64) (ml []*Role, err error) {
// 	tx := db.Set("gorm:auto_preload", true).Begin()
// 	if limit != 0 {
// 		tx = tx.Limit(limit)
// 	} else {
// 		var count int64
// 		tx.Model(&ml).Count(&count)
// 		tx = tx.Limit(count)
// 	}
// 	// Joinsを使用したmany2manyの絞り込みの試し
// 	// tx = tx.
// 	// 	Joins("JOIN role_access_rights ON role_access_rights.role_id = roles.id").
// 	// 	Joins("JOIN access_rights ON role_access_rights.access_right_id = access_rights.id").
// 	// 	Where("access_rights.access_right_name = ?", "MOP")

// 	err = tx.Offset(offset).Find(&ml).Commit().Error
// 	return
// }
