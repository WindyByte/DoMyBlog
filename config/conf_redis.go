package config

type Redis struct {
	Addr       string `yaml:"addr"`        //redis服务器地址，ip:port格式
	Password   string `yaml:"password"`    // redis 密码，默认为空
	DB         int    `yaml:"db"`          // redis DB 数据库，默认为0
	MaxRetries int    `yaml:"max_retries"` // 最大重试次数，默认为3
	PoolSize   int    `yaml:"pool_size"`   // 连接池最大连接数量, 默认为 10 * runtime.GOMAXPROCS
}
