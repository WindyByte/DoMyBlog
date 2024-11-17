package config

type Logger struct {
	LogLevel     string `yaml:"log_level"`
	Prefix       string `yaml:"prefix"`
	Directory    string `yaml:"directory"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}
