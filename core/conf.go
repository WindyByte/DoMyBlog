package core

import (
	"backend/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// InitConfig 读取yaml配置文件
func InitConfig() *config.Config {
	const ConfigFile = "settings.yaml"
	conf := config.Config{}
	yamlByte, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Fatalf("Read Config File Error: %v\n", err)
	}
	err = yaml.Unmarshal(yamlByte, &conf)
	if err != nil {
		log.Fatalf("Unmarshal Config File Error: %v\n", err)
	}
	log.Printf("Init Config: %+v", conf)
	return &conf
}
