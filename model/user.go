package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model `json:"meta"`
	UserID     string `gorm:"index;type:CHAR(36);column:user_id" json:"userId"`               // 主键，指定为 CHAR(36)
	UserName   string `gorm:"type:VARCHAR(100);column:user_name" json:"name,omitempty"`       // 指定为 VARCHAR(100)
	Passwd     string `gorm:"type:TEXT;column:passwd" json:"passwd,omitempty"`                // 指定为 TEXT
	AvatarURL  string `gorm:"type:VARCHAR(255);column:avatar_url" json:"avatarUrl,omitempty"` // 指定为 VARCHAR(255)
	Email      string `gorm:"type:VARCHAR(150);column:email" json:"email,omitempty"`          // 唯一，指定为 VARCHAR(150)
	Role       Role   `gorm:"type:INT;column:role" json:"role,omitempty"`                     // 指定为 ENUM 类型
}

func (UserInfo) TableName() string {
	return "user_infos" // 数据库中的表名
}

func (UserInfo) ParseRole(role Role) string {
	switch role {
	case PermissionAdmin:
		return "admin"
	case PermissionUser:
		return "user"
	case PermissionGuest:
		return "guest"
	default:
		return "disable"
	}
}
