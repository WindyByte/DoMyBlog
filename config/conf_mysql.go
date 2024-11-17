package config

import "strconv"

type MySQL struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Config       string `yaml:"config"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	LogMode      string `yaml:"log_mode"` // 是否开启Gorm全局日志
	LogLevel     string `yaml:"log_level"`
}

// Dsn 数据源名称 (Data Source Name)
func (m MySQL) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Database + "?" + m.Config
}
