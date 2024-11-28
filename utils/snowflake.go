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
	snowflakeClient, err := snowflake.NewNode(1)
	if err != nil {
		core.GetLogger().Errorf("GenID failed, err: %v", err)
		return -1
	}
	return snowflakeClient.Generate().Int64()
}
