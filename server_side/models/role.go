package models

// Role ...
type Role struct {
	Model
	RoleName     string         `gorm:"" json:"roleName"`
	AccessRights []*AccessRight `gorm:"many2many:role_accessRights" json:"accessRights"`
}

func init() {

}

// CreateRole ...
func CreateRole(role Role) (roleID int64, err error) {
	err = db.Create(&role).Error
	return
}

// GetRole ...
func GetRole(roleID int64) (role Role, err error) {
	err = db.Set("gorm:auto_preload", true).First(&role, roleID).Error
	return
}

// GetAllRoles ...
func GetAllRoles(limit int64, offset int64) (ml []*Role, err error) {
	tx := db.Set("gorm:auto_preload", true).Begin()
	if limit != 0 {
		tx = tx.Limit(limit)
	} else {
		var count int64
		tx.Model(&ml).Count(&count)
		tx = tx.Limit(count)
	}

	err = tx.Offset(offset).Find(&ml).Commit().Error
	return
}
