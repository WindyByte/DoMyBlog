package model

import "time"

type Role int

const (
	PermissionAdmin   Role = 1 // 管理员权限
	PermissionUser    Role = 2 // 用户权限
	PermissionGuest   Role = 3 // 游客权限
	PermissionDisable Role = 4 // 禁用权限
)

type MetaInfo struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}
