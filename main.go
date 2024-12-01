package main

import (
	"backend/core"
	"backend/routers"
)

func main() {
	// 初始化配置文件
	r := routers.InitRouter()
	core.GetLogger().Info("Start backend server.")
	err := r.Run(core.GetBasicConfig().System.Addr())
	if err != nil {
		core.GetLogger().Errorf("Start backend server failed, err: %v", err)
	}
	//err := infra.MInsertTestArticleInfo()
	//if err != nil {
	//	core.GetLogger().Errorf("MInsertTestArticleInfo failed, err: %v", err)
	//}
	//err := infra.MInsertTestUserInfo()
	//if err != nil {
	//	core.GetLogger().Errorf("MInsertTestUserInfo failed, err: %v", err)
	//}
	//userInfos, err := infra.MGetUserInfoByUserName("John Doe")
	//if err != nil {
	//	core.GetLogger().Errorf("MGetUserInfoByUserName failed, err: %v", err)
	//}
	//core.GetLogger().Infof("MGetUserInfoByUserName success, userInfos: %v", userInfos)

	//infra.MInsertTestUserInfo()
	//userInfo, err := infra.GetUserInfoByID(1862786145679577089)
	//if err != nil {
	//	core.GetLogger().Errorf("GetUserInfoByID failed, err: %v", err)
	//	return
	//}
	//core.GetLogger().Infof("GetUserInfoByID success, userInfo: %v", userInfo)
	//userInfo.UserName = "wanglei"
	//err = infra.UpdateUserInfoByID(userInfo)
	//if err != nil {
	//	core.GetLogger().Errorf("UpdateUserInfoByID failed, err: %v", err)
	//	return
	//}
	//core.GetLogger().Infof("UpdateUserInfoByID success, userInfo: %v", userInfo)
}
