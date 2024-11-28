package core

import (
	"backend/config"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// initConfig 读取yaml配置文件
func initConfig() *config.BasicConfig {
	const ConfigFile = "settings.yaml"
	conf := config.BasicConfig{}
	yamlByte, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Sprintf("Read Config File Error: %v\n", err))
	}
	err = yaml.Unmarshal(yamlByte, &conf)
	if err != nil {
		panic(fmt.Sprintf("Unmarshal Config File Error: %v\n", err))
	}
	return &conf
}
