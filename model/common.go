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
	ID        uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // 主键，自动递增
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`          // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`          // 更新时间
}

type UserResource struct {
	Name      string `json:"name" msg:"name"`
	EnName    string `json:"enName" msg:"enName"`
	AvatarUrl string `json:"avatarUrl" msg:"avatarUrl"` // nolint: byted_s_variable_name
}
