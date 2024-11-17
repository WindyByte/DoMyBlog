package models

type UserInfo struct {
	ID       int    `gorm:"primaryKey;autoIncrement;comment:'主键'"`       // 主键字段
	Username string `gorm:"type:varchar(50);not null;comment:'用户名'"`     // 用户名字段
	Passwd   string `gorm:"type:varchar(255);not null;comment:'加密后的密码'"` // 密码字段
}
type UserXxxxxx struct {
	ID       int    `gorm:"primaryKey;autoIncrement;comment:'主键'"`       // 主键字段
	Username string `gorm:"type:varchar(50);not null;comment:'用户名'"`     // 用户名字段
	Passwd   string `gorm:"type:varchar(255);not null;comment:'加密后的密码'"` // 密码字段
}
