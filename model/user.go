package model

type UserInfo struct {
	UserID    int64  `gorm:"index;primaryKey;BIGINT;column:user_id" json:"userID"`
	UserName  string `gorm:"type:VARCHAR(100);unique;column:user_name" json:"userName,omitempty"`
	Passwd    string `gorm:"type:VARCHAR(100);column:passwd" json:"passwd,omitempty"`
	AvatarURL string `gorm:"type:VARCHAR(255);column:avatar_url" json:"avatarUrl,omitempty"` // 头像URL
	Email     string `gorm:"type:VARCHAR(150);column:email" json:"email,omitempty"`
	Role      Role   `gorm:"type:TINYINT;column:role" json:"role,omitempty"` // 指定为 TINYINT 类型
	MetaInfo
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
