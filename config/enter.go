package config

type Config struct {
	MySQL  MySQL  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
