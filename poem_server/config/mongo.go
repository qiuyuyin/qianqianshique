package config

type Mongo struct {
	Path        string `mapstructure:"path" json:"path" yaml:"path"`                          // 服务器地址:端口
	Dbname      string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                  // 数据库名
	Username    string `mapstructure:"username" json:"username" yaml:"username"`              // 数据库用户名
	Password    string `mapstructure:"password" json:"password" yaml:"password"`              // 数据库密码
	MaxPoolSize uint64 `mapstructure:"max-pool-size" json:"maxPoolSize" yaml:"max-pool-size"` // 最大连接池大小
	MinPoolSize uint64 `mapstructure:"min-pool-size" json:"minPoolSize" yaml:"min-pool-size"` // 最大连接池大小
}

func (m *Mongo) Uri() string {
	if m.Username == "" || m.Password == "" {
		return "mongodb://" + m.Path
	} else {
		return "mongodb://" + m.Username + ":" + m.Password + "@" + m.Path
	}
}
