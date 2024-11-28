package utils

import (
	"backend/core"
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	snowflakeClient *snowflake.Node
	onceSnowFlake   sync.Once
)

func GenID() int64 {
	if snowflakeClient == nil {
		onceSnowFlake.Do(initSnowflake)
	}
	return snowflakeClient.Generate().Int64()
}

func initSnowflake() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		core.GetLogger().Errorf("initSnowflake, err: %v", err)
		return
	}
	snowflakeClient = node
}
