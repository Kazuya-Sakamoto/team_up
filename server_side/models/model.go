package models

import (
	"time"
)

// Model gorm.Modelにjsonタグを付けて再定義した基底モデル
type Model struct {
	ID        int64      `gorm:"primary_key" json:"id"` // ID
	CreatedAt time.Time  `gorm:"" json:"createdAt"`     // 作成日
	UpdatedAt time.Time  `gorm:"" json:"updatedAt"`     // 更新日
	DeletedAt *time.Time `sql:"index" json:"deletedAt"` // 削除日
}
