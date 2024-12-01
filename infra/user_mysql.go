package infra

import (
	"backend/core"
	"backend/model"
	"backend/utils"
	"errors"
	"gorm.io/gorm"
	"log"
)

func CreateUserInfo(userInfo *model.UserInfo) error {
	db := core.GetGormClient()
	err := db.Create(userInfo).Error
	return err
}

func GetUserInfoByID(UserID int64) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	db := core.GetGormClient()
	result := db.First(&userInfo, UserID)
	return &userInfo, result.Error
}

func MGetUserInfoByID(UserID int64) ([]*model.UserInfo, error) {
	var userInfo []*model.UserInfo
	db := core.GetGormClient()
	result := db.Where("user_id =?", UserID).Find(&userInfo)
	return userInfo, result.Error
}

func GetUserInfoByUserName(UserName string) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	db := core.GetGormClient()
	result := db.Where("user_name =?", UserName).First(&userInfo)
	return &userInfo, result.Error
}

func UpdateUserInfoByID(userInfo *model.UserInfo) error {
	db := core.GetGormClient()
	var existingUser model.UserInfo
	err := db.Where("user_id = ?", userInfo.UserID).First(&existingUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	result := db.Model(&model.UserInfo{}).Where("user_id = ?", userInfo.UserID).Updates(userInfo)
	return result.Error
}

func MGetUserInfoByUserName(UserName string) ([]*model.UserInfo, error) {
	var userInfos []*model.UserInfo
	db := core.GetGormClient()
	result := db.Where("user_name = ?", UserName).Find(&userInfos)
	return userInfos, result.Error
}

func MInsertTestUserInfo() error {
	db := core.GetGormClient()
	testUsers := []model.UserInfo{
		{
			UserID:    utils.GenID(),
			UserName:  "John Doe",
			Passwd:    "hashed_password_123",
			AvatarURL: "https://example.com/avatar1.png",
			Email:     "john.doe@example.com",
			Role:      1,
		},
		{
			UserID:    utils.GenID(),
			UserName:  "Jane Smith",
			Passwd:    "hashed_password_456",
			AvatarURL: "https://example.com/avatar2.png",
			Email:     "jane.smith@example.com",
			Role:      2,
		},
		{
			UserID:    utils.GenID(),
			UserName:  "Alice Johnson",
			Passwd:    "hashed_password_789",
			AvatarURL: "https://example.com/avatar3.png",
			Email:     "alice.johnson@example.com",
			Role:      1,
		},
		{
			UserID:    utils.GenID(),
			UserName:  "John Doe",
			Passwd:    "hashed_password_123",
			AvatarURL: "https://example.com/avatar1.png",
			Email:     "john.doe@example.com",
			Role:      1,
		},
		{
			UserID:    utils.GenID(),
			UserName:  "Jane Smith",
			Passwd:    "hashed_password_456",
			AvatarURL: "https://example.com/avatar2.png",
			Email:     "jane.smith@example.com",
			Role:      2,
		},
		{
			UserID:    utils.GenID(),
			UserName:  "Alice Johnson",
			Passwd:    "hashed_password_789",
			AvatarURL: "https://example.com/avatar3.png",
			Email:     "alice.johnson@example.com",
			Role:      1,
		},
	}
	// 批量插入数据
	if err := db.Create(&testUsers).Error; err != nil {
		return err
	}
	log.Println("Test user info inserted successfully!")
	return nil
}
